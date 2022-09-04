package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcase := []string{"Hello World", " ", "!12345"}
	for _, tc := range testcase {
		f.Add(tc) // 使用 f.Add 提供种子语料库
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			t.Skip()
			// return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			t.Skip()
			// return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
