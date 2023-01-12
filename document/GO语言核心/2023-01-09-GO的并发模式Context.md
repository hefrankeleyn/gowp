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
func Background() Context
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

