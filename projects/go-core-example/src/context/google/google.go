package google

import (
	"context"
	"context/userip"
	"encoding/json"
	"net/http"
)

// 一个Result包含标题和搜索结果的URL
type Result struct {
	Title, URL string
}

type Results []Result

// Search 发送查询到 Google 搜索，并返回结果
func Search(ctx context.Context, query string) (Results, error) {
	// 准备Google 搜索 API 请求
	req, err := http.NewRequest("GET", "https://ajax.googleapis.com/ajax/services/search/web?v=1.0", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Set("q", query)

	// 如果 ctx 懈怠了用户的IP地址，将它转发给服务器。
	// Google APIs 使用用户的IP地址来区分服务器发起的请求和最终用户的请求
	if userIP, ok := userip.FromContext(ctx); ok {
		q.Set("userip", userIP.String())
	}
	req.URL.RawQuery = q.Encode()

	// 发出HTTP请求，并处理响应。
	// 如果 ctx.Done 是被取消， httpDo 函数将取消请求。
	var results Results
	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		// 解析JSON格式的查询结果
		// https://developers.google.com/web-search/docs/#fonje
		var data struct {
			ResponseData struct {
				Results []struct {
					TitleNoFormatting string
					URL               string
				}
			}
		}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return err
		}
		for _, res := range data.ResponseData.Results {
			results = append(results, Result{Title: res.TitleNoFormatting, URL: res.URL})
		}
		return nil
	})
	// httpDo 等待我们提供的闭包返回，所以在这里读取结果是安全的。
	return results, err
}

// httpDo 发起 HTTP 请求， 并使用响应调用 f 。
//  如果 ctx.Done 在请求或f函数运行期间被关闭，httpDo取消那个请求，等待f离开，并返回ctx.Err。
// 否则 返回 f 的 error
func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	// 在一个goroutine中 运行HTTP请求， 并传递响应给 f
	c := make(chan error, 1)
	req = req.WithContext(ctx)
	go func() { c <- f(http.DefaultClient.Do(req)) }()
	select {
	case <-ctx.Done():
		<-c // 为了等待f 函数返回
		return ctx.Err()
	case err := <-c:
		return err
	}
}
