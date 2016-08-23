package protocol

import (
	"bufio"
	"encoding/binary"
	. "fmt"
	"io"
	"net"
	"os"
)

func Read(r io.Reader) ([]byte, error) {
	var bodyLen int32
	binary.Read(r, binary.BigEndian, &bodyLen)
	Println("len", bodyLen)

	body := make([]byte, bodyLen)
	_, err := io.ReadFull(r, body)
	//Println("body", string(body))

	if err != nil {
		return nil, err
	}
	return body, nil
}

func SendResponse(w io.Writer, data []byte) (int, error) {
	err := binary.Write(w, binary.BigEndian, int32(len(data)))
	if err != nil {
		return 0, err
	}

	n, err := w.Write(data)
	if err != nil {
		return 0, err
	}

	return (n + 4), nil
}

func test(err error, mesg string) {
	if err != nil {
		Println("SERVER: ERROR: ", mesg)
		os.Exit(-1)
	} else {
		//Println("SERVER: OK: ", mesg)
	}
}

func main() {
	buf := make([]byte, 1500)

	Println("starting server...")

	netlisten, err := net.Listen("tcp", "127.0.0.1:6666")
	test(err, "Listen")
	defer netlisten.Close()

	for {
		Println("server wait for client ...")
		con, err := netlisten.Accept()
		test(err, "Accept for client")
		Println("client has connect to server")

		writer := bufio.NewWriter(con)
		reader := bufio.NewReader(con)
		reader.Read(buf)

		Println("readed: ", string(buf))
		msg := "response from server"

		writer.WriteString(msg)
		writer.WriteString(msg)
		writer.WriteString(msg)
		writer.Flush()
	}
}
