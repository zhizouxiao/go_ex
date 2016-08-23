package main

import "fmt"
import "time"

func main() {
	before := time.Now().Unix()
	fmt.Println(before)
	t1 := time.NewTimer(time.Second * 2)
	<-t1.C
	fmt.Println("Timer 1 expired!")
	t2 := time.NewTimer(time.Second * 2)
	go func() {
		/* code */
		<-t2.C
		fmt.Println("Timer 2 expired!")
	}()
	stop2 := t2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stoped!")
	}
	end := time.Now().Unix()
	fmt.Println(end)

}
