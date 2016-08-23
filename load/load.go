package main

import "fmt"
import "time"

func main() {
	i := 0
	before := time.Now().Unix()
	fmt.Println(before)
	a := 0
	for i < 10000000000 {
		i = i + 1
		a = a + i
	}
	fmt.Println(time.Now().Unix() - before)
	fmt.Println(a)

}
