package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int) {
	fmt.Println("worker...")
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
	fmt.Println("end worker...")
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	fmt.Println("init worker...")
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	time.Sleep(time.Second * 2)
	fmt.Println("init jobs...")
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	<-results
	for i := 1; i < 10; i++ {
		//<-results
	}

}
