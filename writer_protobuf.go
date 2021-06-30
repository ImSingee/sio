package sio

import (
	"fmt"
)

import "github.com/golang/protobuf/proto"

func (w *Writer) WriteProtoMessage(message proto.Message) (int, error) {
	p, err := proto.Marshal(message)
	if err != nil {
		return 0, fmt.Errorf("cannot marshal proto message: %w", err)
	}

	// Write VarUInt length
	x, err := w.WriteVarUInt(uint64(len(p)))
	if err != nil {
		return x, fmt.Errorf("cannot write length of proto message: %w", err)
	}

	// Write message
	y, err := w.Write(p)
	if err != nil {
		return x + y, fmt.Errorf("cannot write proto message data: %w", err)
	}

	return x + y, nil
}
