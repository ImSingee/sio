package sio

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

func (r *Reader) ReadProtoMessage(message proto.Message) error {
	length, err := r.ReadVarUInt()
	if err != nil {
		return fmt.Errorf("cannot read proto message length: %w", err)
	}

	p, err := r.ReadBytes(int(length))
	if err != nil {
		return fmt.Errorf("cannot read proto message data: %w", err)
	}

	err = proto.Unmarshal(p, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal proto message: %w", err)
	}

	return nil
}
