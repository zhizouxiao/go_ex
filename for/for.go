package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	for {
		i := 0
		if i > 10 {
			break
		}
		fmt.Print(i)
		i = i + 1
		break
	}

}
