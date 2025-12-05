package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for a := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, a)
		time.Sleep(time.Second)
		results <- a * 2
	}
}
func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	// start 3 workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}
	// send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	close(results)
	for r := 1; r <= 5; r++ {
		fmt.Println("Result: ", <-results)
	}
}
