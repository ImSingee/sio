package sio

import (
	"bytes"
	"io"
)

func Equal(one_, another_ io.Reader) (bool, error) {
	const chunkSize = 40960 // 40 KB

	one := NewReader(one_)
	another := NewReader(another_)

	for {
		b1 := make([]byte, chunkSize)
		_, err1 := one.Read(b1)

		b2 := make([]byte, chunkSize)
		_, err2 := another.Read(b2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true, nil
			} else if err1 == io.EOF || err2 == io.EOF {
				return false, nil
			} else if err1 != nil {
				return false, err1
			} else if err2 != nil {
				return false, err2
			}
		}

		if !bytes.Equal(b1, b2) {
			return false, nil
		}
	}
}
