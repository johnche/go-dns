package dns

import (
	"encoding/binary"
	"math"
	"math/rand"
	"strings"
)

func uint16ToBytes(u uint16) []byte {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, u)
	return bytes
}

func randUint16() uint16 {
	return uint16(rand.Intn(math.MaxUint16))
}

func encodeDomain(domainName string) []byte {
	var encodedDomain []byte
	for _, level := range strings.Split(domainName, ".") {
		encodedDomain = append(encodedDomain, byte(len(level)))
		encodedDomain = append(encodedDomain, []byte(level)...)
	}

	// null terminator
	encodedDomain = append(encodedDomain, byte(0))

	return encodedDomain
}
