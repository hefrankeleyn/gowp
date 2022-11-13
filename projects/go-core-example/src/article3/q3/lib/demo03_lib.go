package lib // 为了不让代码包的使用者产生困惑，我们总是应该让声明的包名与其父目录的名称一致

import (
	in "article3/q3/lib/internal"
	"os"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
