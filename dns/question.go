package dns

const (
	CLASS_IN uint16 = 1
	TYPE_A   uint16 = 1
)

type Question struct {
	Name  string
	Type  uint16
	Class uint16
}

func (q Question) Bytes() []byte {
	var data []byte
	data = append(data, encodeDomain(q.Name)...)
	data = append(data, uint16ToBytes(q.Type)...)
	data = append(data, uint16ToBytes(q.Class)...)

	return data
}

func ParseQuestion(data []byte) (Question, int) {
	bytePointer := 0

	domainName, length := decodeDomain(data[12:])
	bytePointer += length

	questionType := bytesToUint16(data[bytePointer : bytePointer+2])
	bytePointer += 2

	questionClass := bytesToUint16(data[bytePointer : bytePointer+2])
	bytePointer += 2

	responseQuestion := Question{
		Name:  domainName,
		Type:  questionType,
		Class: questionClass,
	}

	return responseQuestion, bytePointer
}
