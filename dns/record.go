package dns

type Record struct {
	Name  []byte
	Type  uint16
	Class uint16
	TTL   uint16
	Data  []byte
}

// dns compression
// https://datatracker.ietf.org/doc/html/rfc1035#section-4.1.4

func ParseRecord(data []byte) Record {
	return Record{}
}
