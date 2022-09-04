package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcase := []string{"Hello World", " ", "!12345"}
	for _, tc := range testcase {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
