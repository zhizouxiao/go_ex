package main

import "fmt"

func main() {
	i := 21
	p := &i
	fmt.Println(*p)
	*p = 44
	fmt.Println(*p)
}
