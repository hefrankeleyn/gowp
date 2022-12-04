package main

import "fmt"

func main() {
	var m map[string]int
	// m = map[string]int{} // 补充上这句话，下面的执行就不会引起恐慌了
	fmt.Println(m)

	key := "one"
	v01, ok := m[key] // 不会引起恐慌
	fmt.Printf("在值为nil的map中获取key为 %q 的值：%d,%v\n", key, v01, ok)
	fmt.Printf("map的长度为:%d\n", len(m))
	fmt.Printf("在值为nil的map上删除key为 %q\n", key)
	delete(m, key) // 不会引起恐慌
	fmt.Println("向值为nil的map上添加一个元素，会引起恐慌")
	m[key] = 2
}
