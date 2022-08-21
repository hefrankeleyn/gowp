package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName 调用 greetings.Hello 带有一个name， 检查有效的返回值
func TestHelloName(t *testing.T) {
	name := "hef"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("hef") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty 调用Hello， 传递一个空字符串，检查一个error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
