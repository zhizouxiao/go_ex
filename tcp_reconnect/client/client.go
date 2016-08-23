package main

import (
	"bufio"
	. "fmt"
	"net"
	"os"
	"time"
)

func test(err error, msg string) {
	if err != nil {
		Println("CLIENT: ERROR: ", msg, err)
		os.Exit(-1)
	} else {
		//Println("CLIENT: OK: ", msg)
	}
}

type Reconnect struct {
	conn net.Conn
}

func (reconnect *Reconnect) dial() {
	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		Println(t)
		con, err2 := net.Dial("tcp", "127.0.0.1:6666")
		if err2 == nil {
			ticker.Stop()
			Println("client has connected to server")
			reconnect.conn = con
			return
		}
	}
}

func main() {
	buf := make([]byte, 1500)

	//con, err := net.Dial("tcp", "127.0.0.1:6666")
	//test(err, "dialing")
	reconnect := new(Reconnect)
	reconnect.dial()

	for {
		con := reconnect.conn
		reader := bufio.NewReader(con)
		writer := bufio.NewWriter(con)
		writer.WriteString("B: hi!\n")
		writer.Flush()

		_, err := reader.Read(buf)
		if err != nil {
			reconnect.dial()
			//return
		}
		Println(string(buf))

	}
	defer reconnect.conn.Close()
}
