package dns

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

const (
	RECURSION_DESIRED uint16 = 1 << 8
)

// header format
// https://datatracker.ietf.org/doc/html/rfc1035#section-4.1.1
type Header struct {
	ID             uint16
	Flags          uint16
	QuestionCount  uint16
	AnswerCount    uint16
	AuthorityCount uint16
	Additionals    uint16
}

func (h Header) Bytes() []byte {
	// Using reflection and usafe calls to create byte array directly from struct
	data := make([]byte, 0, int(unsafe.Sizeof(Header{})))
	buffer := bytes.NewBuffer(data)

	buffer.Reset()
	binary.Write(buffer, binary.BigEndian, h)

	return buffer.Bytes()
}

func ParseHeader(data []byte) Header {
	return Header{
		ID:             binary.BigEndian.Uint16(data[0:2]),
		Flags:          binary.BigEndian.Uint16(data[2:4]),
		QuestionCount:  binary.BigEndian.Uint16(data[4:6]),
		AnswerCount:    binary.BigEndian.Uint16(data[6:8]),
		AuthorityCount: binary.BigEndian.Uint16(data[8:10]),
		Additionals:    binary.BigEndian.Uint16(data[10:12]),
	}
}
