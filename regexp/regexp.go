package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	wordRx := regexp.MustCompile("[A-Za-z]+")
	line := "%assd%as%dasdsa"
	line = wordRx.ReplaceAllString(line, "game")
	fmt.Println(line)
	dir, _ := filepath.Split(os.Args[0])
	fmt.Println(dir)

}
