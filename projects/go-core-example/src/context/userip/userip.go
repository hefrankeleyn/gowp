package userip

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

func FromRequest(req *http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}
	return userIP, nil
}

// 密钥类型未导出以防止与其他包中定义的上下文密钥发生冲突。
type key int

// userIPKey 是一个context key，用于用户的IP地址。它的零值是随意定的。
// 如果这个包定义了其它的context key。它们应该是不同的数值。
const userIPKey key = 0

// NewContext 返回一个新的Context， 携带userIP
func NewContext(ctx context.Context, userIP net.IP) context.Context {
	return context.WithValue(ctx, userIPKey, userIP)
}

// 如果ip地址存在， FromContext 从ctx中提取用户的ip地址，
func FromContext(ctx context.Context) (net.IP, bool) {
	// 如果 ctx 没有针对key的值，ctx.Value 将返回 nil
	// net.IP 类型的断言 将对 nil 返回 ok=false
	userIP, ok := ctx.Value(userIPKey).(net.IP)
	return userIP, ok
}
