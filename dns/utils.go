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

func bytesToUint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
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

func decodeDomain(data []byte) (string, int) {
	domainName := ""
	bytePointer := 0
	for {
		// byte is an alias for uint8
		length := int(data[bytePointer])
		bytePointer++

		domainName += string(data[bytePointer : bytePointer+length])
		bytePointer += length

		// null terminator upcoming
		if data[bytePointer] == 0x00 {
			break
		}

		// nothing more to read
		if bytePointer >= len(data) {
			break
		}

		// theres another segment upcoming, lets add dot between
		domainName += "."
	}

	return domainName, bytePointer
}
