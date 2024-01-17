package dns

const (
	CLASS_IN uint16 = 1
	TYPE_A            uint16 = 1
)

type Question struct {
	Name  []byte
	Type  uint16
	Class uint16
}

func (q Question) Bytes() []byte {
	var data []byte
	data = append(data, q.Name...)
	data = append(data, uint16ToBytes(q.Type)...)
	data = append(data, uint16ToBytes(q.Class)...)

	return data
}

type QuestionData []byte
func (data QuestionData) toQuestion() Question {
	domainName := ""
	bytePointer := 0
	for {
		// byte -> uint8 -> int, necessary?
		length := int(uint8(data[bytePointer]))
		bytePointer++

		domainName += string(data[bytePointer:bytePointer + length])
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

	return Question{}
}
