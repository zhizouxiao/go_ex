package main

import (
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("write_test", "ch3", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %v", string(message.Body))
		//wg.Done()
		return nil
	}))
	//err := q.ConnectToNSQD("127.0.0.1:4150")
	//err := q.ConnectToNSQD("127.0.0.1:4250")
	//err := q.ConnectToNSQLookupd("127.0.0.1:4161")
	err := q.ConnectToNSQDs([]string{"127.0.0.1:4150", "111.203.187.159:4150"})
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()

}
