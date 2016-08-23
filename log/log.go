package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: debug ", log.Lshortfile)
	logger.Print("Hello, log file!")
	logger.Error("Hello, log file!")

	fmt.Print(&buf)
}
