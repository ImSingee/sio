package sio

import (
	"encoding/binary"
	"io"
	"unsafe"
)

type Writer struct {
	writer io.Writer
}

func NewWriter(wt io.Writer) *Writer {
	b, ok := wt.(*Writer)

	if ok {
		return b
	}

	return &Writer{wt}
}

func (w *Writer) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

func (w *Writer) WriteByte(b byte) error {
	_, err := w.writer.Write([]byte{b})

	return err
}

// 写入全部字符串，如果无法全部写入依然会返回错误
func (w *Writer) WriteString(s string) (int, error) {
	return w.writer.Write(*(*[]byte)(unsafe.Pointer(&s)))
}

func (w *Writer) WriteUInt8(b uint8) (int, error) {
	return w.writer.Write([]byte{b})
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

//
//// 注意：n 的数量不能大于 buffer size
//func (r *Reader) Peek(n int) ([]byte, error) {
//	return r.breader.Peek(n)
//}
//
//// 尽可能多的跳过 n 个字节
//func (r *Reader) Skip(n int) (int, error) {
//	discard, err := r.breader.Discard(n)
//
//	if err == nil { // 跳过足够
//		return discard, err
//	}
//
//	if err == io.EOF && discard != n { // 未跳过足够
//		return discard, nil
//	}
//
//	// 其他异常
//	return discard, err
//}
//
//func (r *Reader) MustSkip(n int) (int, error) {
//	return r.breader.Discard(n)
//}
//
//func (r *Reader) ReadByte() (byte, error) {
//	return r.breader.ReadByte()
//}
//
//func (r *Reader) ReadBytes(n int) ([]byte, error) {
//	if n < 0 {
//		return nil, fmt.Errorf("cannot read negative bytes")
//	}
//	if n == 0 {
//		return nil, nil
//	}
//
//	result := make([]byte, n)
//
//	read, err := r.Read(result)
//	if err != nil {
//		return result[:read], err
//	}
//
//	return result[:read], nil
//}
//
//// 读取足够的字节，如果不够会报错
//// 当 n <= buffer size 时，该函数保证如果读不够则回退
//// 否则，会停留在最后的读取地
//func (r *Reader) ReadEnoughBytes(n int) ([]byte, error) {
//	// 如果 buffer 足够，利用 peek 完成完美回退
//	bytes, err := r.breader.Peek(n)
//	if err != bufio.ErrBufferFull {
//		if err != nil {
//			return nil, err
//		}
//		if len(bytes) != n {
//			return bytes, fmt.Errorf("no enough bytes can be read")
//		}
//
//		return bytes, nil
//	}
//
//	// 否则正常读取，不考虑回退（TODO 优化）
//	bytes = make([]byte, n)
//	read, err := io.ReadFull(r, bytes)
//
//	// ReadFull 保证读不够会返回错误，因此可以不额外判断
//	return bytes[:read], err
//}
//
//func (r *Reader) ReadBytesAsString(n int) (string, error) {
//	bytes, err := r.ReadEnoughBytes(n)
//
//	if err != nil {
//		return "", err
//	}
//
//	return *((*string)(unsafe.Pointer(&bytes))), nil
//}
//
//func (r *Reader) ReadUInt8() (uint8, error) {
//	return r.ReadByte()
//}
//
//func (r *Reader) ReadUInt16() (uint16, error) {
//	bytes, err := r.ReadEnoughBytes(2)
//
//	if err != nil {
//		return 0, err
//	}
//
//	return binary.BigEndian.Uint16(bytes), err
//}
//
//func (r *Reader) ReadUInt32() (uint32, error) {
//	bytes, err := r.ReadEnoughBytes(4)
//
//	if err != nil {
//		return 0, err
//	}
//
//	return binary.BigEndian.Uint32(bytes), err
//}
//
//func (r *Reader) ReadUInt64() (uint64, error) {
//	bytes, err := r.ReadEnoughBytes(8)
//
//	if err != nil {
//		return 0, err
//	}
//
//	return binary.BigEndian.Uint64(bytes), err
//}
