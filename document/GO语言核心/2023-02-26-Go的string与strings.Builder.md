# Go的string与strings.Builder

[toc]



## 一、strings.Builder 的优势

- 以存在的内容不可变，但可以拼接更多的内容；
- 减少了内存分配和内存拷贝的次数；
- 可将可容重叠，可重用值；

## 二、string类型的值

在Go语言中，string类型的值是不可变的。

基于原字符串的裁剪和拼接操作：

- 裁剪操作可以使用切片表达式；
- 拼接操作可以使用操作符+实现；

一个string的内容会被存储在一块连续的内存空间中，这个连续空间的字节数量也会被记录下来，并用于表示该string值的长度。

**对string类型的值，执行len() 方法，得到的是字节的数量。**

```go
	var str01 string
	str01 = "Go，你好"
	// len(string) : 11
	fmt.Printf("len(string) : %v \n", len(str01))
	// len([]char): 5
	fmt.Printf("len([]char): %v\n", len([]rune(str01)))
	// len([]byte)： 11
	fmt.Printf("len([]byte)： %v\n", len([]byte(str01)))
```

**对string类型的值，进行切片操作，相当于对底层的字节数组做切片**

```go
	// 在string类型的值上进行切片操作，相当于对底层的字节数组做切片
	fmt.Println(str01[0:3]) // Go�
	fmt.Println(str01[0:5]) // Go，
```

在执行字符串拼接的时候，会把所有被拼接的字符串依次拷贝到一个崭新且足够大的连续内存空间中，并把相应指针指的string值作为结果返回。

## 三、与string相比，Builder的优势体现在拼接方面

与string相比，strings.Builder 的优势最要体现在字符串拼接方面。

Builder有一个内容容器，它是一个以byte为元素类型的切片（简称字节切片）。

**Builder类型的值只能被拼接或完全覆盖。**

### 3.1 Builder的拼接，与Builder的自动扩容

可以通过Write、WriteByte、WriteRune、WriteString方法，把新的内容拼接到已存在的内容的尾部。

Builder会自动地对自身的内容容器进行扩容。扩容策略与切片的扩容策略一致。

### 3.2 手动扩容

Builder的值还可以手动扩容，通过调用Builder的Grow方法，就可以做到。

Grow方法接受一个int类型的参数n，这个参数表示要扩充的字节数量。

如果未用容量大于或等于n，Grow方法可能什么都不做。

### 3.3 Builder 的重用

通过调用Reset方法，可以让Builder值重新回到零值的状态，就像它从未被使用一样。

## 四、strings.Builder的使用约束

- springs.Builder 类型的值在已被真正使用后，就不可再被复制。
- 由于其内容不是完全不可变的，需要使用方自行解决操作冲突和并发安全的问题。

只要调用了Builder值的拼接或扩容方法，就意味着真正使用。一旦真正使用，Builder值就不能被复制，否则调用副本的方法会引起panic。

```go
	var sb01 strings.Builder
	sb01.WriteString("Go")
	sb02 := sb01
	// sb02.Grow(1) // 这里会引起恐慌panic
	_ = sb02
```

**虽然Builder值不可复制，但是它的指针指却可以复制。复制的指针值会指向同一个Builder值。**

**对于处于零值状态的Builder值，复制不会有任何问题**。所以，只要在传递之前调用Reset方法即可。

```go
	var sb01 strings.Builder
	sb01.WriteString("Go")
  sb01.Reset()
	sb03 := sb01
	sb03.Grow(1)
```

## 五、strings.Reader 类型的值可以高效读取字符串

strings.Reader 类型的值，可以让我们很方便的读取一个字符串中的内容。因为，在读取的过程中，Reader值会保存已读取的字节的计数。

**一读计数代表着下次读取的起始索引位置。**

> Reader 正是依靠这样的一个计数，以及针对字符串的切片表达式，从而实现快读读取。

通过Len方法和Size方法获取已读计数：

```go
	var reader1 strings.Reader
	reader1 = *strings.NewReader("Go， 你好")
	ch, size, err := reader1.ReadRune()
	fmt.Printf("%v, %v, %v\n", string(ch), size, err)
	readingIndex := reader1.Size() - int64(reader1.Len())
	fmt.Printf("以读计数：%v\n", readingIndex)
```

Reader  的Seek方法用于设定下一次读取的起始索引位置；如果把io.SeekCurrent 的值作为第二个参数值传递给该方法。那么它会依据当前的已读计数，以及第一个参数offset的值来计算新的计数值。

