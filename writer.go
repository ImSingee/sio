package sio

import (
	"encoding/binary"
	"io"
	"unsafe"
)

type Writer struct {
	writer io.Writer
	n      uint64
}

func NewWriter(wt io.Writer) *Writer {
	b, ok := wt.(*Writer)

	if ok {
		return b
	}

	return &Writer{writer: wt}
}

// N 返回内部写计数器
func (w *Writer) N() uint64 {
	return w.n
}

// ResetN 将内部写计数器清零
func (w *Writer) ResetN() {
	w.n = 0
}

func (w *Writer) Write(p []byte) (n int, err error) {
	n, err = w.writer.Write(p)
	w.n += uint64(n)
	return
}

func (w *Writer) WriteByte(b byte) error {
	n, err := w.writer.Write([]byte{b})
	w.n += uint64(n)

	return err
}

// WriteString 写入全部字符串，如果无法全部写入依然会返回错误
func (w *Writer) WriteString(s string) (n int, err error) {
	n, err = w.writer.Write(*(*[]byte)(unsafe.Pointer(&s)))
	w.n += uint64(n)

	return
}

func (w *Writer) WriteUInt8(b uint8) (int, error) {
	return w.Write([]byte{b})
}

func (w *Writer) WriteUInt16(x uint16) (int, error) {
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, x)

	return w.Write(bs)
}

func (w *Writer) WriteUInt32(x uint32) (int, error) {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, x)

	return w.Write(bs)
}

func (w *Writer) WriteUInt64(x uint64) (int, error) {
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, x)

	return w.Write(bs)
}

func (w *Writer) WriteZeros(n int) (int, error) {
	bs := make([]byte, n)

	return w.Write(bs)
}

func (w *Writer) WriteVarUInt(x uint64) (int, error) {
	buf := make([]byte, 16) // 10 byte is enough
	n := binary.PutUvarint(buf, x)
	return w.Write(buf[:n])
}

func (w *Writer) WriteVarInt(x int64) (int, error) {
	buf := make([]byte, 16) // 10 byte is enough
	n := binary.PutVarint(buf, x)
	return w.Write(buf[:n])
}
