// package lib2 // 第一个改动， 把 main 改为 lib2。（故意让声明的包名和所在的目录名称不同）
package lib // 为了不让代码包的使用者产生困惑，我们总是应该让声明的包名与其父目录的名称一致

import (
	"fmt"
)

// 第二个改动， 把首字母小写的 hello 改为 首字母大写的 Hello
func Hello(name string) {
	fmt.Println(name)
}
