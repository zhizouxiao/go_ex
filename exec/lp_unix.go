package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	//"strings"
)

func main() {
	cmd := exec.Command("date")
	//cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("in all caps: ", out.String())
}
