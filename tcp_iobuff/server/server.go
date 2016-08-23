package main

import (
	"bufio"
	. "fmt"
	protocol "github.com/zhizouxiao/go_ex/tcp_iobuff"
	"net"
	"os"
)

func test(err error, mesg string) {
	if err != nil {
		Println("SERVER: ERROR: ", mesg)
		os.Exit(-1)
	} else {
		//Println("SERVER: OK: ", mesg)
	}
}

func main() {
	//buf := make([]byte, 1500)

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
		//reader := bufio.NewReader(con)
		//reader.Read(buf)
		data, err := protocol.Read(con)

		Println("readed: ", string(data))
		msg := "response from server"
		protocol.SendResponse(writer, []byte(msg))
		protocol.SendResponse(writer, []byte(msg))
		protocol.SendResponse(writer, []byte(msg))

		//writer.WriteString(msg)
		//writer.WriteString(msg)
		//writer.WriteString(msg)
		writer.Flush()
	}
}
