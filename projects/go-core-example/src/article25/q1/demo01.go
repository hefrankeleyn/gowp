package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

// bufPool 代表存放数据缓冲区的临时对象池
var bufPool sync.Pool

// Buffer 代表一个简单的数据块缓冲区的接口
type Buffer interface {
	// Delimiter 用于获取数据块之间的定界符
	Delimiter() byte
	// write 用于写一个数据块
	Write(contents string) (err error)
	// Read 用于读一个数据块
	Read() (contexts string, err error)
	// Free 用于释放当前的缓冲区
	Free()
}

// myBuffer 代表了数据块缓冲区的一种实现
type myBuffer struct {
	buf       bytes.Buffer
	delimiter byte
}

func (b *myBuffer) Delimiter() byte {
	return b.delimiter
}

func (b *myBuffer) Write(contexts string) (err error) {
	if _, err = b.buf.WriteString(contexts); err != nil {
		return
	}
	return b.buf.WriteByte(b.delimiter)
}

func (b *myBuffer) Read() (contexts string, err error) {
	return b.buf.ReadString(b.delimiter)
}

func (b *myBuffer) Free() {
	bufPool.Put(b)
}

// delimiter 代表预定义的定界符
var delimiter = byte('\n')

// GetBuffer 用于获取一个数据块缓冲区
func GetBuffer() Buffer {
	return bufPool.Get().(Buffer)
}

func init() {
	bufPool = sync.Pool{
		New: func() any {
			return &myBuffer{delimiter: delimiter}
		},
	}
}

func main() {
	buf := GetBuffer()
	defer buf.Free()
	// buf := &myBuffer{delimiter: delimiter}
	buf.Write("一个池是一个临时对象集合，这些对象能够单独的保存和释放。")
	buf.Write("一个池让多个goroutine同时使用也是安全的。")
	buf.Write("一个池在第一次被使用之后，不能被拷贝")

	fmt.Println("数据块在缓冲区中")
	for {
		block, err := buf.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Errorf("未知的错误：%s", err))
		}
		fmt.Print(block)
	}
}
