package main

import (
	"context"
	"context/google"
	"context/userip"
	"log"
	"net/http"
	"text/template"
	"time"
)

func main() {
	// 注册 handleSearch 来处理 /search 端点
	http.HandleFunc("/search", handleSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))
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
	// 运行Google 搜索，并打印结果
	start := time.Now()
	results, err := google.Search(ctx, query)
	elapsed := time.Since(start)
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

}

var resultsTemplate = template.Must(template.New("results").Parse(`
<html>
<head/>
<body>
<ol>
{{range .Results}}
<li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a></li>
{{end}}
</ol>
<p>{{len .Results}} result in {{.Elapsed}}; timeout {{.Timeout}}</p>
</body>
</html>
`))
