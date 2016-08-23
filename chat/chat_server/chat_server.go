package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type Client struct {
	incoming chan string
	outgoing chan string
	reader   *bufio.Reader
	writer   *bufio.Writer
	name     string
}

func (client *Client) Read() {
	for {
		line, err := client.reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		client.incoming <- "cli" + client.name + ":" + line
	}
}

func (client *Client) Write() {
	for data := range client.outgoing {
		client.writer.WriteString(data + "\n")
		client.writer.Flush()
	}
}

func (client *Client) Listen() {
	go client.Read()
	go client.Write()
}

func NewClient(connection net.Conn, name string) *Client {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)

	client := &Client{
		incoming: make(chan string),
		outgoing: make(chan string),
		reader:   reader,
		writer:   writer,
		name:     name,
	}

	client.Listen()

	return client
}

type ChatRoom struct {
	clients  []*Client
	joins    chan net.Conn
	incoming chan string
	outgoing chan string
}

func (chatRoom *ChatRoom) Broadcast(data string) {
	for _, client := range chatRoom.clients {
		client.outgoing <- data
	}
}

func (chatRoom *ChatRoom) Join(connection net.Conn) {
	client := NewClient(connection, fmt.Sprint(len(chatRoom.clients)+1))
	chatRoom.clients = append(chatRoom.clients, client)
	go func() {
		for {
			chatRoom.incoming <- <-client.incoming
		}
	}()
}

func (chatRoom *ChatRoom) Listen() {
	go func() {
		for {
			select {
			case data := <-chatRoom.incoming:
				chatRoom.Broadcast(data)
			case conn := <-chatRoom.joins:
				chatRoom.Join(conn)
			}
		}
	}()
}

func NewChatRoom() *ChatRoom {
	chatRoom := &ChatRoom{
		clients:  make([]*Client, 0),
		joins:    make(chan net.Conn),
		incoming: make(chan string),
		outgoing: make(chan string),
	}

	chatRoom.Listen()

	return chatRoom
}

func main() {
	chatRoom := NewChatRoom()

	listener, _ := net.Listen("tcp", ":6666")
	id := 1

	for {
		conn, _ := listener.Accept()
		fmt.Println("new client comming...", id)
		id = id + 1
		chatRoom.joins <- conn

	}
}
