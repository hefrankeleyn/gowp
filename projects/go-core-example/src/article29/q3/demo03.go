package main

import (
	"bytes"
	"io"
	"strings"
)

func main() {
	var b *bytes.Buffer
	b = bytes.NewBufferString("ab")
	_ = interface{}(b).(io.ByteReader)
	_ = interface{}(b).(io.RuneReader)
	_ = interface{}(b).(io.ByteScanner)
	_ = interface{}(b).(io.RuneScanner)
	// io.ByteReader
	var reader01 *strings.Reader
	reader01 = strings.NewReader("aa")
	_ = interface{}(reader01).(io.ByteReader)
	_ = interface{}(reader01).(io.RuneReader)
	_ = interface{}(reader01).(io.ByteScanner)
	_ = interface{}(reader01).(io.RuneScanner)
	//
	var reader02 *strings.Reader
	reader02 = strings.NewReader("aa")
	_ = interface{}(reader02).(io.Seeker)
	_ = interface{}(reader02).(io.ReaderAt)
	var sectionReader01 *io.SectionReader
	sectionReader01 = io.NewSectionReader(reader02, 0, 1)
	_ = interface{}(sectionReader01).(io.Seeker)
	_ = interface{}(sectionReader01).(io.ReaderAt)
}
