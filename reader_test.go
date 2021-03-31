package sio

import (
	"bytes"
	"github.com/ImSingee/tt"
	"testing"
)

func TestReadBytesAsUInt16(t *testing.T) {
	// 0x0102 = 258
	reader := NewReader(bytes.NewReader([]byte{1, 2}))

	uInt16, err := reader.ReadUInt16()
	tt.AssertIsNil(t, err)
	tt.AssertEqual(t, uint16(258), uInt16)
}

func TestWriteBytesAsUInt16(t *testing.T) {
	p
}
