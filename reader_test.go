package sio

import (
	"bytes"
	"strings"
	"testing"

	"github.com/ImSingee/tt"
)

func TestReadBytesAsUInt16(t *testing.T) {
	// 0x0102 = 258
	reader := NewReader(bytes.NewReader([]byte{1, 2}))

	uInt16, err := reader.ReadUInt16()
	tt.AssertIsNil(t, err)
	tt.AssertEqual(t, uint16(258), uInt16)
}

func TestWriteBytesAsUInt16(t *testing.T) {
	// TODO
}

func TestReadLine(t *testing.T) {
	r := NewReader(strings.NewReader("hello"))
	result, eof, err := r.ReadLine()
	tt.AssertIsNil(t, err)
	tt.AssertEqual(t, "hello", result)
	tt.AssertEqual(t, true, eof)

	r = NewReader(strings.NewReader("hello\n"))
	result, eof, err = r.ReadLine()
	tt.AssertIsNil(t, err)
	tt.AssertEqual(t, "hello", result)
	tt.AssertEqual(t, false, eof)
	result, eof, err = r.ReadLine()
	tt.AssertIsNil(t, err)
	tt.AssertEqual(t, "", result)
	tt.AssertEqual(t, true, eof)

	r = NewReader(strings.NewReader("hello\nworld"))
	result, eof, err = r.ReadLine()
	tt.AssertIsNil(t, err)
	tt.AssertEqual(t, "hello", result)
	tt.AssertEqual(t, false, eof)
	result, eof, err = r.ReadLine()
	tt.AssertIsNil(t, err)
	tt.AssertEqual(t, "world", result)
	tt.AssertEqual(t, true, eof)
}
