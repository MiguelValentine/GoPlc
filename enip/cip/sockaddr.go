package cip

import (
	"bytes"
	"encoding/binary"
	"github.com/MiguelValentine/goplc/enip/etype"
)

// big endian
type SocketAddress struct {
	TypeID    etype.XUINT
	Length    etype.XUINT
	SinFamily etype.XUINT
	SinPort   etype.XUINT
	SinAddr   etype.XUDINT
	SinZero   uint64
}

func (s *SocketAddress) Buffer() []byte {
	buffer := new(bytes.Buffer)
	_ = binary.Write(buffer, binary.LittleEndian, s.TypeID)
	_ = binary.Write(buffer, binary.LittleEndian, s.Length)
	_ = binary.Write(buffer, binary.BigEndian, s.SinFamily)
	_ = binary.Write(buffer, binary.BigEndian, s.SinPort)
	_ = binary.Write(buffer, binary.BigEndian, s.SinAddr)
	_ = binary.Write(buffer, binary.BigEndian, s.SinZero)
	return buffer.Bytes()
}