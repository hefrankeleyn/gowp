# GO的并发模式Context

[toc]

## 一、介绍

参考：[Go的并发模式Context](https://go.dev/blog/context)。

在Go服务中，每一个传入进来的请求都将在它自己的goroutine中被处理。请求处理器通常启动额外的goroutines去访问后端（例如数据库和RPC服务）。基于请求进行工作的goroutines通常需要获取请求的特殊值，例如最终的用户身份、授权令牌、和请求的最后期限。当一个请求被取消或者超时，所有给予请求进行工作的goroutines应该立刻退出，以便于系统能够回收任何它们使用的资源。

在Google中，我们开发了一个`context`包，以便于更容易的传递请求范围的值、取消信号、和最终的读取API边界的期限，涉及处理请求中的所有goroutines。这个包是公开可用的，被称为`context`。这篇文章将详细描述如何使用这个包，并提供完成的工作示例。

## 二、Context

`context`包的核心代码是`Context`类型：

```go
// 一个Context懈怠了终止日期、取消信号、和请求范围的值，读取API的边界。
// 它的方法对于多个goroutines同时使用是安全的。
type Context interface{
  // Done方法，返回一个通道。当Context被取消或者超时的时候，这个通道将被关闭。
  Done() <-chan struct{}
  // Err 用于表明为什么Context是被取消。通常是由于Done的通道被关闭
  Err() error
  // Deadline 当无论什么原因，这个context被取消的时候，返回一个时间，
  Deadline() (deadline time.Time, ok bool)
  // Value 返回与key对应的值，如果没有就返回nil
  Value(key interface{}) interface{}
}
```

Done方法返回一个通道，这个通道给那些代表`Context`运行函数的取消信号：当通道被关闭，当通道关闭，这些函数应该立刻放弃它们的工作并返回。Err函数返回一个错误，表明为什么Context 是被取消。在[管道和消除](https://go.dev/blog/pipelines)这篇文件中更详细的探讨了Done 通道的完整用法。

一个Context没有一个Cancel方法，和Done通道仅用于接收的原因相同：函数通常是接收到一个取消信号，而不是发送一个信号。尤其是，当一个父操作为子操作启动goroutines，这些子操作不应该能够去取消父操作。相反，`WithCancel`函数（在下面描述）提供了一个方法用于取消一个新的Context值。

一个Context被多个goroutine同时使用时安全的。代码能够传统同一个Context给任何数量的goroutines，并取消Context以向所有goroutines发出信号。

Deadline方法允许函数去确定是否它们应该都开始工作；如果剩余的时间太少，它可能是没有价值的。代码也可以使用最后期限给I/O操作设置超时时间。

Value允许一个Context可以携带请求范围的值。这个数据对多个goroutines同时使用必须是安全的。

## 三、context的衍生

context包提供了一个函数用于从已经存在的Contexts里衍生新的Context值。**这些值形成了树**：当一个Context被取消，所有从它衍生出来的Context也将被取消。

Background是任何Context树的根，它永远不会被取消：

```go
/// Background 返回一个空的context。它绝不会被取消，没有最后期限，不带值。
//  Background 通常在main、init、test中使用，作为请求顶层的Context。
func Background() Context
```

`WithCancel`和`WithTimeout`返回衍生的Context值。它能比它父Context更早的取消。

关联一个传入请求的Context，当请求处理返回的时候，它通常能够被取消。

当使用多个副本的时候，`WithCancel`对于·取消冗余的请求也很有用。

`WithTimeout`对于设置后台服务请求的截止日期很有用。

```go
// WithCancel 返回一个父Context的拷贝，它的Done通道在父级关闭时立即关闭。
// Done 完成关闭或取消
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
// 一个CancelFunc 取消一个Context
type CancelFunc func()

// WithTimeout 返回一个父级的副本，它的Done 通道在父级关闭时立即关闭。
// 关闭 Done、调用取消或超时。
// 新Context的最后期限，是现在+超时时间 和 父的截止日期（如果有的话）中较早的一个。
// 如果计时器仍在运行，取消函数将释放其资源。
func WithTimeout(parent Context, timeout time.Duration) (context, CancelFunc)
```

`WithContext` 提供了一个方法，去将请求范围的值和Context进行关联。

```go
// WithValue 返回一个父Context的副本， 它的Value方法返回key对应的val。
func WithValue(parent Context, key interface{}, value interface{}) Context 
```

最好查看如何使用context包的方法是通过工作示例。

## 四、示例：Google Web Search

我们的示例是一个HTTP服务，它处理像`/search?q=golang&timeout=1s`这样的URL，通过转发"golang" 的查询，到[Google Web Search API](https://developers.google.com/custom-search?hl=zh-cn)并展示结果。这个`timeout`参数告诉服务端在一段时间后取消请求。

代码被分到三个包中：

- `server` 提供main函数，并处理`/search`请求；
- `userip` 提供用于从请求中提取用户ip地址的函数，并将它和Context进行关联；
- `google` 提供查询函数用于发送请求到 Google；

### 4.1 server程序

`server`程序处理像`/search?q=golang`的请求，通过提供前几个对golang在Google上的查询结果。它注册handleSearch来处理`/search`端点。这个处理创建了一个初始的Context称之为ctx，当处理返回的时候，安排它取消。如果请求包含`timeout`URL参数，当超时时间通过Context将自动被取消：

```go
func main() {
	// 注册 handleSearch 来处理 /search 端点
	http.HandleFunc("/search", handleSearch)
}

func handleSearch(w http.ResponseWriter, req *http.Request) {
	// ctx 是 这个处理器的 Context。
	// 调用 cancel 关闭ctx.Done 的通道。这是对这个请求的取消信号，被处理器启动。
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	// 获取超时时间
	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	if err == nil {
		//  请求有超时时间，因此创建一个context，当超时时间到期后它将自动取消
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	// handleSearch返回之后，立刻取消ctx
	defer cancel()
```

处理器从请求中提取查询，并通过调用userip包提取客户端的IP地址。客户端的IP地址对后端是需要的，因此handleSearch将它绑定到ctx上：

```go
	// 检查查询请求
	query := req.FormValue("q")
	if query == "" {
		http.Error(w, "no query", http.StatusBadRequest)
		return
	}
	userIP, err := userip.FromRequest(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx = userip.NewContext(ctx, userIP)
```

handle调用`google.Search`方法，带有ctx和query：

```go
	// 运行Google 搜索，并打印结果
	start := time.Now()
	results, err := google.Search(ctx, query)
	elapsed := time.Since(start)
```

如果查询成功，处理器渲染结果

```go
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := resultsTemplate.Execute(w, struct {
		Results          google.Results
		Timeout, Elapsed time.Duration
	}{
		Results: results,
		Timeout: timeout,
		Elapsed: elapsed,
	}); err != nil {
		log.Print(err)
		return
	}
```

### 4.2 userip 包

`userip`包提供了用于从请求中提取用户ip地址的函数，并将它关联到Context中。一个Context提供了键-值对的映射，在这里键和值都是`interface{}`类型。键的类型必须支持判等操作，值在被多个goroutine同时使用多时候必须是安全的。像userip包，隐藏了这个映射的细节，并提供了抢类型用于读取一个特定的Context值。

为了避免键的冲突，`userip`定义了一个不对外开放的类型key，并使用这个类型的值作为context的键。

```go
// 密钥类型未导出以防止与其他包中定义的上下文密钥发生冲突。
type key int

// userIPKey 是一个context key，用于用户的IP地址。它的零值是随意定的。
// 如果这个包定义了其它的context key。它们应该是不同的数值。
const userIPKey key = 0
```

`FromRequest`从一个`http.Request`中提取一个userIP：

```go
func FromRequest(req *http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}
```

`NewContext`返回一个新的Context，携带了一个提供的userIP值。

```go
// 如果ip地址存在， FromContext 从ctx中提取用户的ip地址，
func FromContext(ctx context.Context) (net.IP, bool) {
	// 如果 ctx 没有针对key的值，ctx.Value 将返回 nil
	// net.IP 类型的断言 将对 nil 返回 ok=false
	userIP, ok := ctx.Value(userIPKey).(net.IP)
	return userIP, ok
}
```

### 4.3 google 包

`google.Search`函数创建一个HTTP请求给`Google Web Search API`，并解析JSON编码的结果。它接受一个Context参数，在请求被处理期间，如果`ctx.Done`被关闭，它将立即返回。

`Google Web Search API`请求，包含了查询请求，并使用IP作为请求参数。

```go
// Search 发送查询到 Google 搜索，并返回结果
func Search(ctx context.Context, query string) (Results, error) {
	// 准备Google 搜索 API 请求
	req, err := http.NewRequest("GET", "https://ajax.googleapis.com/ajax/services/search/web?v=1.0", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Set("q", query)

	// 如果 ctx 懈怠了用户的IP地址，将它转发给服务器。
	// Google APIs 使用用户的IP地址来区分服务器发起的请求和最终用户的请求
	if userIP, ok := userip.FromContext(ctx); ok {
		q.Set("userip", userIP.String())
	}
	req.URL.RawQuery = q.Encode()
```

`Search`使用一个有用的函数，httpDo，用于发出http请求，并在请求或响应在处理期间，如果ctx.Done 被关闭，那么就取消它们。`Search`将一个闭包传递给httpDo，用于处理http响应。

```go
	// 发出HTTP请求，并处理响应。
	// 如果 ctx.Done 是被取消， httpDo 函数将取消请求。
	var results Results
	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		// 解析JSON格式的查询结果
		// https://developers.google.com/web-search/docs/#fonje
		var data struct {
			ResponseData struct {
				Results []struct {
					TitleNoFormatting string
					URL               string
				}
			}
		}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return err
		}
		for _, res := range data.ResponseData.Results {
			results = append(results, Result{Title: res.TitleNoFormatting, URL: res.URL})
		}
		return nil
	})
	// httpDo 等待我们提供的闭包返回，所以在这里读取结果是安全的。
	return results, err
```

httpDo 函数返回http请求，并在一个新的goroutine中处理它的响应。在goroutine离开之前，如果`ctx.Done`被关闭，它将取消请求。

```go
// httpDo 发起 HTTP 请求， 并使用响应调用 f 。
//  如果 ctx.Done 在请求或f函数运行期间被关闭，httpDo取消那个请求，等待f离开，并返回ctx.Err。
// 否则 返回 f 的 error
func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	// 在一个goroutine中 运行HTTP请求， 并传递响应给 f
	c := make(chan error, 1)
	req = req.WithContext(ctx)
	go func() { c <- f(http.DefaultClient.Do(req)) }()
	select {
	case <-ctx.Done():
		<-c // 为了等待f 函数返回
		return ctx.Err()
	case err := <-c:
		return err
	}
}
```

## 五、使用context包中程序实体实现`sync.WaitGroup`同样的功能

### （1）使用`sync.WaitGroup`实现一对多goroutine协作流程多同步工具

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	coordinateWithWaitGroup()
}

func coordinateWithWaitGroup() {
	total := 12
	stride := 3
	var num int32
	fmt.Printf("The number: %d [with sync.WaitGroup]\n", num)
	var wg sync.WaitGroup
	for i := 1; i <= total; i += stride {
		wg.Add(stride)
		for j := 0; j < stride; j++ {
			go addNum(&num, i+j, wg.Done)
		}
		wg.Wait()
	}
}

func addNum(numP *int32, id int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	for i := 0; ; i++ {
		currNum := atomic.LoadInt32(numP)
		newNum := currNum + 1
		if atomic.CompareAndSwapInt32(numP, currNum, newNum) {
			fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
			break
		} else {
			fmt.Printf("The CAS option failed. [%d-%d]\n", id, i)
		}
	}
}

```

### （2）使用context包中程序实体来实现

```go
func coordinateWithContext() {
	total := 12
	var num int32
	fmt.Printf("The number: %d [with context.Context]\n", num)
	cxt, cancelFunc := context.WithCancel(context.Background())
	for i := 1; i <= total; i++ {
		go addNum(&num, i, func() {
			// 如果所有的addNum函数都执行完毕，那么就立即分发子任务的goroutine
			// 这里分发子任务的goroutine，就是执行 coordinateWithContext 函数的goroutine.
			if atomic.LoadInt32(&num) == int32(total) {
				// <-cxt.Done() 针对该函数返回的通道进行接收操作。
				// cancelFunc() 函数被调用，针对该通道的接收会马上结束。
				// 所以，这样做就可以实现“等待所有的addNum函数都执行完毕”的功能
				cancelFunc()
			}
		})
	}
	<-cxt.Done()
	fmt.Println("end.")
}
```

```shell
$ go run demo01.go
The number: 0 [with context.Context]
The number: 1 [12-0]
The number: 3 [6-0]
The number: 4 [7-0]
The number: 5 [8-0]
The number: 6 [9-0]
The number: 2 [5-0]
The number: 8 [10-0]
The number: 10 [11-0]
The number: 11 [1-0]
The number: 9 [3-0]
The number: 7 [2-0]
end.
```

执行发现，有时候结果并不对。

## 六、Context的特点

### （1）Context介绍：上下文、树、根节点

Context类型是一种非常通用的同步工具，它的值不但可以任意扩散，还可以被用来传递额外的信息和信号。

> Context类型可以提供一类代表上下文的值。此类值是并发安全的，也就是说它可以被传播给多个goroutine。

Context类型的值可以繁衍，这意味着可以通过一个Context值产生任意个子值。这些子值可以携带其父值的属性和数据，也可以响应我们通过父值传递的信号。

所有的Context值共同构成了一颗代表上下文全貌的树形结构。这棵树的树根（或称上下文的根节点）是一个已经在context包中预定义好的Context值，它是全局唯一的。通过`context.Background`函数可以取到它。

> 上下文根节点仅仅是一个最基本的支点，它不提供任何额外的功能。也就是说它既不可以被撤销，也不能携带任何数据。

### （2）四个延伸Context值的函数

在context包中，包含四个用于繁衍Context值的函数，即WithCancel 、WithDeadline、WithTimeout、WithValue。

- 四个函数的第一个参数的类型都是`context.Context`，名称都为parent。

  > 顾名思义，这个位置上的参数对应的都是它们将会产生的Context 值的父值。

- WithCancel用于产生一个可撤销的parent的子值；

   通过调用该函数，可以得到一个衍生自上下文根节点的Context值，和一个用于发送撤销信号的函数。

  ```go
  cxt, cancelFunc := context.WithCancel(context.Background())
  ```

- `WithDeadline`和 `WithTimeout`函数都是用来产生一个会定时撤销的parent的子值。

- `WithValue`函数可以用来产生一个会携带额外数据的parent的子值。

### （3）理解context包中“可撤销”和“撤销一个context”

在Context接口中，有两个与撤销息息相关的方法：

- `Done`方法会返回一个元素类型为`struct{}`的接收通道。（让使用方感知撤销信号）

  这个接收通道的用途并不是传递元素值，而是让调用方法感知“撤销”当前Context值的那个信号。

  一旦当前的Contex值被撤销，这里的接收通道会立刻关闭。对于一个未包含任何元素值的通道来说，它的关闭会使任何针对它的接收操作立即结束。

- `Err`方法。让使用方得到撤销的具体原因。

  Context的Err方法的结果是error类型，并且其值值可能等于`context.Canceled`变量的值，或者`context.DeadlineExceeded` 变量的值。

  `context.Canceled` 表示手动撤销；`context.DeadlineExceeded` 表示由于给定的过期时间已到，而导致的撤销。

对“撤销“和”可撤销“对理解：

- 如果把”撤销“当作名词理解：指的是用来表达“撤销”状态的信号；
- 如果把“撤销”当作动词理解：指的是对撤销信号的表达；
- “可撤销”：指的是具有传达这种撤销信号的能力；

理解`context.WithCancel`：

- 通过调用`context.WithCancel`可以产生一个可撤销的Context值，还会获得一个用于触发撤销信号的函数。

- 通过调用这个用于触发撤销信号的函数， 撤销信号会被传达给这个Context值，并由它的Done方法的结果值（一个接收通道）表达出来；

- 撤销函数只负责触发信号，而对应的可撤销的Context值也只负责传达信号。它们都不会去管后面具体的“撤销”操作。

  实际上，我们的代码在感知到撤销操作以后，可以任意的操作，Context对此并没有任何的约束。

### （4）撤销信号在上下文树中的传播

在contex包中，包含四个用于繁衍Context值的函数，其中`WithCancel`、`WithDeadline`、`WithTimeout`都是被用于基于给定的Context值产生可撤销的子值的。

> `context.WithValue`函数得到的Context值可不撤销，撤销信号在被传播时，若遇到它们则会直接跨过，并试图将信号直接传给它们的子值。

context  包的WithCancel函数在被调用后会产生两个结果值，第一个结果值就是那个可撤销的Context值，而第二个则是用于触发撤销信号的函数。

撤销函数被调用之后，对应的Context值会先关闭它内部的接收通道，也就是它的Done方法会返回的那个通道。然后，它会向它的所有子值（或者说子节点）传达撤销信号。这些子值会如法炮制，把撤销信号继续传播下去。最后这个Context值会断开它与其父值之间的关联。

通过调用context包中的`WithDealline`函数，或者`WithTimeout`函数生成的Context值也是可撤销的。它们不但可以被手动撤销，还会依据在生成时被给定的过期时间，自动地进行撤销。

> 这里的定时撤销，是借助内部的计时器来做。

当过期时间到达时，`WithDeadline`和`WithTimeout`这两个Context值当行为与Context值被手动撤销时的行为几乎是一致的。只不过`WithDeadline`和`WithTimeout`会在最后停止并释放掉其内部的计时器。

### （5）通过Context携带数据

`WithValue`函数在产生新的Context值的时候需要三个参数：父值、键、值。

> 与字典对于键的约束类似，这里键的类型必须是可判等的。原因很简单，我们从中获取数据的时候，它需要根据给定的键来查找对应的值。只不过，这种Context值并不是用字典来存储键和值的，只是简单的存储在响应的字段中。

Context类型的Value方法就是用来获取数据的。在我们调用含数据的Context值的Value方法时，它会先判断给定的键，是否与当前值中存储的键相等。如果相等就把该值中的存储的值直接返回，否则就到父值中继续寻找。如果父值中仍然未存储相等的键，那么该方法就会沿着上下文根节点的方向一路查找下去。

注意，除了含数据的Context值以外，其它几种Context值都无法携带数据。因此，Context值的Value方法在沿路查找的时候，会直接跨过那几种值。

Context接口并没有提供改变数据的方法。因此，在通常情况下，我们只能通过在上下文树中添加含数据的Context值来存储新的数据，或者通过撤销此种值的父值丢弃掉相应的数据。如果你存储在这里的数据可以从外部改变，那么必须自行保证安全。



