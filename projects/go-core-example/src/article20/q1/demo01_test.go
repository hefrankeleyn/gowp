package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	var name string
	greeting, err := hello(name)
	if err == nil {
		t.Errorf("err 是 nil，但是它不应该如此，(name=%q)", name)
	}
	if greeting != "" {
		t.Errorf("greeting 不是空，但是它应该为空，(name=%q)", name)
	}
	name = "Robert"
	greeting, err = hello(name)
	if err != nil {
		t.Errorf("err 不是 nil，但是它应该为nil，(name=%q)", name)
	}
	if greeting == "" {
		t.Errorf("greeting 是空，但是它不应该为空，(name=%q)", name)
	}
	expected := fmt.Sprintf("hello, %s!", name)
	if greeting != expected {
		t.Errorf("实际值 %q 和期待的值不一样， (name = %q)", greeting, name)
	}
	t.Logf("期待的值为%q\n", expected)
}

func TestIntroduce(t *testing.T) {
	into := introduce()
	excepted := "Welcome to my Golang column."
	if into != excepted {
		t.Errorf("introduce 的实际值%q和期望的值不一样", into)
	}
	t.Logf("introduce 期望的值为 %q\n", excepted)
}

// func TestFail(t *testing.T) {
// 	// t.Fail()
// 	t.FailNow() // 次调用会让当前的测试立刻失效
// 	t.Log("Fail.")
// }

// func TestError(t *testing.T) {
// 	t.Error("t.Error 相当于t.Log，再调用 t.Fail")
// 	t.Errorf("t.Error 相当于%s，再调用 t.Fail", "t.logf")
// }

func TestFatal(t *testing.T) {
	// t.Fatal("t.Error 相当于t.Log，再调用 t.FailNew")
	t.Fatalf("t.Error 相当于%s，再调用 t.FailNew", "t.logf")
}
