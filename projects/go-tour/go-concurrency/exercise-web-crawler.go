package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// 根据URL返回内容，和在页面中发现的URLs切片
	Fetch(url string) (body string, urls []string, err error)
}

type SafeURLMap struct {
	mu sync.Mutex
	mb map[string]bool
}

func (m *SafeURLMap) GetVal(key string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.mb[key]
}

func (m *SafeURLMap) SetVal(key string, val bool) {
	m.mu.Lock()
	m.mb[key] = val
	m.mu.Unlock()
}

// 获取并放默认值，放在一个锁里
func (m *SafeURLMap) PutTrueIfAbsent(key string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	res := m.mb[key]
	m.mb[key] = true
	return res
}

// Crawl 使用fetcher 去递归爬取页面，从一个url开始，并到达最大深度
func Crawl(url string, depth int, fetcher Fetcher, c chan bool, m *SafeURLMap) {
	// TODO: 并发的获取URL
	// TODO：不获取相同的URL两次
	// 判断 该URL是否已经获取，如果已经获取，直接结束
	// if ok := m.GetVal(url); ok {
	// 	c <- true
	// 	return
	// }
	// 设置URL
	// m.SetVal(url, true)
	if ok := m.PutTrueIfAbsent(url); ok {
		// 已经获取过了
		c <- true
		return
	}
	// 已经达到某一个深度，直接返回
	if depth <= 0 {
		c <- false
		return
	}
	body, urls, err := fetcher.Fetch(url)
	// URL 错误，直接返回
	if err != nil {
		fmt.Println(err)
		c <- false
		return
	}
	fmt.Printf("found:%s %q\n", url, body)
	c2l := make([]chan bool, len(urls))
	for i, u := range urls {
		c2l[i] = make(chan bool)
		go Crawl(u, depth-1, fetcher, c2l[i], m)
	}
	for _, v := range c2l {
		<-v
	}
	c <- true
	return
}

func main() {
	// 创建一个缓存，用于避免二次访问
	m := &SafeURLMap{mb: make(map[string]bool)}
	c := make(chan bool)
	go Crawl("https://golang.org/", 4, fetcher, c, m)
	// 阻塞，等待程序结束
	<-c
}

type fakeResult struct {
	body string
	urls []string
}

// fakeFetcher 是 一个 Fetcher，获取预设的结果
type fakeFetcher map[string]*fakeResult

func (f fakeFetcher) Fetch(url string) (body string, urls []string, err error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("没有发现：%s", url)
}

// fetcher 是一个填充的fakeFetcher
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
