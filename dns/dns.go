package dns

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Record struct {
	Name  []byte
	Type  uint16
	Class uint16
	TTL   uint16
	Data  []byte
}

func CreateQuery(domainName string, recordType uint16) []byte {
	header := Header{
		ID:             randUint16(),
		Flags:          RECURSION_DESIRED,
		QuestionCount:  1,
		AnswerCount:    0,
		AuthorityCount: 0,
		Additionals:    0,
	}

	question := Question{
		Name:  encodeDomain(domainName),
		Type:  recordType,
		Class: CLASS_IN,
	}
	fmt.Printf("header: %x\n", header.Bytes())
	fmt.Printf("question: %x\n", question.Bytes())

	var query []byte
	query = append(query, header.Bytes()...)
	query = append(query, question.Bytes()...)

	// DEBUG
	query = append(query, byte('\n'))

	return query
}

func Client() {
	// port 53 is dns port
	//conn, err := net.Dial("udp", "8.8.8.8:53")
	conn, err := net.Dial("udp", "localhost:5050")
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	query := CreateQuery("www.example.com", TYPE_A)
	fmt.Printf("created query: %x\n", query)

	if _, err := conn.Write(query); err != nil {
		log.Fatalf("failed to send query: %v", err)
	}

	response := make([]byte, 1024)
	if _, err = bufio.NewReader(conn).Read(response); err != nil {
		log.Fatalf("failed reading response: %v", err)
	}

	fmt.Printf("response header: %x\n", response[0:12])
	fmt.Printf("response body: %x\n", response[12:])
	//responseHeader := HeaderData(response[0:12]).toHeader()
}
