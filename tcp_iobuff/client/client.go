package main

import (
	"bufio"
	. "fmt"
	protocol "github.com/zhizouxiao/go_ex/tcp_iobuff"
	"net"
	"os"
)

func test(err error, msg string) {
	if err != nil {
		Println("CLIENT: ERROR: ", msg, err)
		os.Exit(-1)
	} else {
		//Println("CLIENT: OK: ", msg)
	}
}

func main() {
	//buf := make([]byte, 1500)

	con, err := net.Dial("tcp", "127.0.0.1:6666")
	test(err, "dialing")
	defer con.Close()
	Println("client has connected to server")

	//reader := bufio.NewReader(con)
	writer := bufio.NewWriter(con)
	msg := "B: hi!"
	protocol.SendResponse(writer, []byte(msg))
	writer.WriteString("B: hi!\n")
	writer.Flush()

	//reader.Read(buf)
	var result []byte
	result, err = protocol.Read(con)
	Println(string(result))
	result, err = protocol.Read(con)
	Println(string(result))
	result, err = protocol.Read(con)
	Println(string(result))
	result, err = protocol.Read(con)
	Println(string(result))
}
