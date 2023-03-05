package main

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

func executeIfNoErr(err error, f func()) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	f()
}

func example1(comment string) {
	// 创建一个字符串
	// 创建一个字符串读取器，它的名字是 reader1。
	fmt.Println("创建一个字符串读取器，它的名字是 reader1。")
	reader1 := strings.NewReader(comment)
	buf1 := make([]byte, 7)
	n, err := reader1.Read(buf1)
	var index1, offset1 int64
	executeIfNoErr(err, func() {
		// Read(7): "Package"
		fmt.Printf("Read(%d): %q\n", n, buf1[:n])
		offset1 = int64(54)
		index1, err = reader1.Seek(offset1, io.SeekCurrent)
	})
	executeIfNoErr(err, func() {
		fmt.Printf("基于当前的所以，移动%d的偏移量后，新的索引值为: %d \n", offset1, index1)
		n, err = reader1.Read(buf1)
	})
	executeIfNoErr(err, func() {
		fmt.Printf("Read(%d):%q\n", n, buf1[:n])
	})
	fmt.Println()
}

func example2(comment string) {
	reader1 := strings.NewReader(comment)
	reader1.Reset(comment)
	num := int64(7)
	fmt.Printf("创建一个新的限制Reader，限制读的数量为：%d\n", num)
	reader2 := io.LimitReader(reader1, num)
	buf2 := make([]byte, 10)
	for i := 0; i < 3; i++ {
		n, err := reader2.Read(buf2)
		executeIfNoErr(err, func() {
			fmt.Printf("Read(%d):%q\n", n, buf2[:n])
		})
	}
	fmt.Println()
}

func example3(comment string) {
	reader1 := strings.NewReader(comment)
	writer1 := new(strings.Builder)
	fmt.Println("创建一个新的teeReader， 带有一个reader和一个writer")
	reader3 := io.TeeReader(reader1, writer1)
	buf4 := make([]byte, 40)
	for i := 0; i < 8; i++ {
		n, err := reader3.Read(buf4)
		executeIfNoErr(err, func() {
			fmt.Printf("Read(%d):%q\n", n, buf4[:n])
		})
	}
	fmt.Println()
}

func example4(comment string) {
	reader1 := strings.NewReader(comment)
	offset1 := int64(56)
	num2 := int64(72)
	fmt.Printf("创建一个section Reader 带有一个Reader， 偏移量为%d, 数量为 %d...\n", offset1, num2)
	reader2 := io.NewSectionReader(reader1, offset1, num2)
	buf1 := make([]byte, 20)
	for i := 0; i < 5; i++ {
		n, err := reader2.Read(buf1)
		executeIfNoErr(err, func() {
			fmt.Printf("Read(%d): %q\n", n, buf1[:n])
		})
	}
	fmt.Println()
}

func example5() {
	reader01 := strings.NewReader("MultiReader returns a Reader that's the logical concatenation of " +
		"the provided input readers.")
	reader02 := strings.NewReader("They're read sequentially.")
	reader03 := strings.NewReader("Once all inputs have returned EOF, " +
		"Read will return EOF.")
	reader04 := strings.NewReader("If any of the readers return a non-nil, " +
		"non-EOF error, Read will return that error.")
	fmt.Println("创建一个multi-reader， 带有4个reader")
	reader1 := io.MultiReader(reader01, reader02, reader03, reader04)
	buf2 := make([]byte, 50)
	for i := 0; i < 8; i++ {
		n, err := reader1.Read(buf2)
		executeIfNoErr(err, func() {
			fmt.Printf("Read(%d): %q\n", n, buf2[:n])
		})
	}
	fmt.Println()
}

func example6() {
	fmt.Println("创建一个新的内存同步管道....")
	pipeReader, pipWriter := io.Pipe()
	_ = interface{}(pipeReader).(io.ReadCloser)
	_ = interface{}(pipWriter).(io.WriteCloser)
	comments := [][]byte{
		[]byte("Pipe creates a synchronous in-memory pipe."),
		[]byte("It can be used to connect code expecting an io.Reader "),
		[]byte("with code expecting an io.Writer."),
	}
	// 这里的同步工具，纯属为了保证下面示例中的打印语句能够执行完成
	// 在实际中没必要这样做
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, d := range comments {
			time.Sleep(time.Millisecond * 500)
			n, err := pipWriter.Write(d)
			if err != nil {
				fmt.Printf("read error : %v\n", err)
				break
			}
			fmt.Printf("Writen(%d): %q\n", n, d)
		}
		pipWriter.Close()
	}()
	go func() {
		defer wg.Done()
		wBuf := make([]byte, 55)
		for {
			n, err := pipeReader.Read(wBuf)
			if err != nil {
				fmt.Printf("read error: %v\n", err)
				break
			}
			fmt.Printf("Read(%d): %q\n", n, wBuf[:n])
		}
	}()
	wg.Wait()
}

func main() {
	comment := "Package io provides basic interfaces to I/O primitives. " +
		"Its primary job is to wrap existing implementations of such primitives, " +
		"such as those in package os, " +
		"into shared public interfaces that abstract the functionality, " +
		"plus some other related primitives."
	// 示例1:: Seek
	example1(comment)
	// 示例2: LimitReader
	example2(comment)
	// 示例3: TeeReader
	example3(comment)
	// 示例4: NewSectionReader
	example4(comment)
	// 示例5: MultiReader
	example5()
	// 示例6
	example6()
}
