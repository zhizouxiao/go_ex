package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

func main() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	for {
		err := w.Publish("write_test", []byte("test4250"))
		if err != nil {
			log.Panic("Could not connect")
		}
		time.Sleep(time.Second)
	}

	w.Stop()
}
