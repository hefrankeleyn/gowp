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
func CancelFunc func()

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

## 一、使用context包中程序实体实现`sync.WaitGroup`同样的功能

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

