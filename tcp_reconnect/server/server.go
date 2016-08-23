package main

import (
	"bufio"
	. "fmt"
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

		writer.WriteString("response from server")
		writer.Flush()
	}
}
