package sio

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"unsafe"
)

var DefaultBufSize int = 4096

type Reader struct {
	breader *bufio.Reader
	n       uint64
}

func NewReader(rd io.Reader) *Reader {
	return NewReaderSize(rd, DefaultBufSize)
}

// 创建一个新的 Reader
// 如果 buf 不存在
func NewReaderSize(rd io.Reader, size int) *Reader {
	b, ok := rd.(*Reader)

	if size <= 64 {
		size = 64
	}

	if ok && b.breader.Size() >= size {
		return b
	}

	r := new(Reader)
	r.breader = bufio.NewReaderSize(rd, size)

	return r
}

// N 返回内部读计数器
func (r *Reader) N() uint64 {
	return r.n
}

// ResetN 将内部读计数器清零
func (r *Reader) ResetN() {
	r.n = 0
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.breader.Read(p)
	r.n += uint64(n)
	return n, err
}

// 注意：n 的数量不能大于 buffer size
func (r *Reader) Peek(n int) ([]byte, error) {
	return r.breader.Peek(n)
}

// 尽可能多的跳过 n 个字节
func (r *Reader) Skip(n int) (int, error) {
	discard, err := r.breader.Discard(n)
	r.n += uint64(discard)

	if err == nil { // 跳过足够
		return discard, err
	}

	if err == io.EOF && discard != n { // 未跳过足够
		return discard, nil
	}

	// 其他异常
	return discard, err
}

func (r *Reader) MustSkip(n int) (int, error) {
	discard, err := r.breader.Discard(n)
	r.n += uint64(discard)
	return discard, err
}

func (r *Reader) ReadByte() (byte, error) {
	b, err := r.breader.ReadByte()
	if err == nil {
		r.n += 1
	}

	return b, err
}

func (r *Reader) ReadBytes(n int) ([]byte, error) {
	if n < 0 {
		return nil, fmt.Errorf("cannot read negative bytes")
	}
	if n == 0 {
		return nil, nil
	}

	result := make([]byte, n)

	read, err := r.Read(result)
	if err != nil {
		return result[:read], err
	}
	r.n += uint64(read)

	return result[:read], nil
}

// 读取足够的字节，如果不够会报错
// 当 n <= buffer size 时，该函数保证如果读不够则回退
// 否则，会停留在最后的读取地
func (r *Reader) ReadEnoughBytes(n int) ([]byte, error) {
	// 如果 buffer 足够，利用 peek 完成完美回退
	bytes, err := r.breader.Peek(n)
	if err != bufio.ErrBufferFull {
		if err != nil {
			return nil, err
		}
		if len(bytes) != n {
			return bytes, fmt.Errorf("no enough bytes can be read")
		}

		_, _ = r.breader.Discard(n)
		r.n += uint64(n)

		return bytes, nil
	}

	// 否则正常读取，不考虑回退（TODO 优化）
	bytes = make([]byte, n)
	read, err := io.ReadFull(r, bytes)

	// ReadFull 保证读不够会返回错误，因此可以不额外判断
	r.n += uint64(read)
	return bytes[:read], err
}

func (r *Reader) ReadBytesAsString(n int) (string, error) {
	bytes, err := r.ReadEnoughBytes(n)

	if err != nil {
		return "", err
	}

	return *((*string)(unsafe.Pointer(&bytes))), nil
}

func (r *Reader) ReadUInt8() (uint8, error) {
	return r.ReadByte()
}

func (r *Reader) ReadUInt16() (uint16, error) {
	bytes, err := r.ReadEnoughBytes(2)

	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint16(bytes), err
}

func (r *Reader) ReadUInt32() (uint32, error) {
	bytes, err := r.ReadEnoughBytes(4)

	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint32(bytes), err
}

func (r *Reader) ReadUInt64() (uint64, error) {
	bytes, err := r.ReadEnoughBytes(8)

	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint64(bytes), err
}

func (r *Reader) ReadVarUInt() (uint64, error) {
	// 内部调用了 ReadByte
	return binary.ReadUvarint(r)
}

func (r *Reader) ReadVarInt() (int64, error) {
	// 内部调用了 ReadByte
	return binary.ReadVarint(r)
}
