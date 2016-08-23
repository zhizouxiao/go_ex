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
		Println("ERROR: ", msg)
		os.Exit(-1)
	} else {
		//Println("CLIENT: OK: ", msg)
	}
}

func main() {
	if len(os.Args) <= 1 {
		Println("please input chat name before chat...")
		return
	}
	chatName := os.Args[1]

	con, err := net.Dial("tcp", "127.0.0.1:6666")
	test(err, "server dead!")
	defer con.Close()
	Println("client has connected to server")
	reader := bufio.NewReader(con)
	con.SetReadDeadline(time.Now().Add(2 * time.Second))

	go func(reader *bufio.Reader) {
		for {
			line, _ := reader.ReadString('\n')
			Println(line)
		}
	}(reader)

	for {
		var input string
		Scanln(&input)

		input = chatName + ":" + input
		con.Write([]byte(input + "\n"))

	}
}
