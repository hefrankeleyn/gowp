# goroutine 的使用

[toc]

## 一、怎么才能让主goroutine等待其它goroutine

### （1）方法一：让主goroutine "sleep"一段时间

```go
func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Microsecond * 100)
}
```

缺点是：睡眠的时间很难把握。

### （2）方法二：使用通道，使用`chan struct{}`类型

```go
func main() {
	num := 10
	var ch = make(chan struct{}, num)
	for i := 0; i < num; i++ {
		go func(i int) {
			fmt.Println(i)
			ch <- struct{}{}
		}(i)
	}
	for i := 0; i < num; i++ {
		<-ch
	}
}
```

类型字面量`struct{}`有些类似空接口类型`interface{}`，它代表了即不包含任何字段也不拥有任何方法的空结构体类型。

**`struct{}`类型值的表示法只有一个，即：`struct{}{}`。并且它占用的内存空间是0字节，确切地说，这个值在整个GO程序中永远都只会存在一份。**

### （3）方法三：使用`sync.WaitGroup`

待补充。

## 二、怎么让多个goroutine按照既定的顺序运行

```go
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	num := uint32(100)
	var count uint32 = 0
	trigger := func(i uint32, fn func()) {
		for {
			if atomic.LoadUint32(&count) == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			// 这里加Sleep语句是很有必要的
			time.Sleep(time.Microsecond)
		}
	}
	for i := uint32(0); i < num; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	trigger(num, func() {})
}
```

这里的trigger函数实现了一种自旋（spining）。

上面的自旋中添加了`time.Sleep(time.Microsecond)`语句：

> 这主要是因为：Go 调度器在需要的时候只会对正在运行的 goroutine 发出通知，试图让它停下来。但是，它却不会也不能强行让一个 goroutine 停下来。 
>
> 所以，如果一条 for 语句过于简单的话，比如这里的 for 语句就很简单（因为里面只有一条 if 语句），那么当前的 goroutine 就可能不会去正常响应（或者说没有机会响应）Go 调度器的停止通知。 
>
> 因此，这里加一个 sleep 是为了：在任何情况下（如任何版本的 Go、任何计算平台下的 Go、任何的 CPU 核心数等），内含这条 for 语句的这些 goroutine 都能够正常地响应停止通知。

不加Sleep语句，可能会导致一直抢占不到资源，也就没有机会运行，就可能会导致程序一直运行，不会终止。

```
乐观锁：总是假设在“我”操作共享资源的过程中没有“其他人”竞争操作。如果发现“其他人”确实在此期间竞争了，也就是发现假设失败，那就等一等再操作。CAS原子操作基本上能够体现出这种思想。通常，低频的并发操作适合用乐观锁。乐观锁一般会用比较轻量级的同步方法（如原子操作），但也不是100%。注意，高频的操作用乐观锁的话反而有可能影响性能，因为多了一步“探查是否有人与我竞争”的操作（当然了，标准的CAS操作可以把这种影响降到最低）。

悲观锁：总是假设在“我”操作共享资源的过程中一定有“其他人”竞争操作。所以“我”会先用某种同步方法（如互斥锁）保护我的操作。这样的话，“我”在将要操作的时候就没必要去探查是否有人与我竞争（因为“我”总是假设肯定有竞争，而且已经做好了保护）。通常，频次较高的并发操作适合用悲观锁。不过，如果并发操作的频次非常低，用悲观锁也是可以的，因为这种情况下对性能影响不大。

最后，一定要注意，使用任何同步方法和异步方法都首先要考虑程序的正确性，并且还要考虑程序的性能。程序的正确性一定要靠功能测试来保障，程序的性能一定要靠性能测试来保障。
```

