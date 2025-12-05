package main

import (
	"fmt"
	"sync"
)

func printEven(oddChan, evenChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i = i + 2 {
		<-evenChan
		fmt.Println(i)
		if i < 10 {
			oddChan <- true
		}
	}
}
func printOdd(oddChan, evenChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i = i + 2 {
		<-oddChan
		fmt.Println(i)
		evenChan <- true
	}
}
func main() {
	var wg sync.WaitGroup
	oddChan := make(chan bool)
	evenChan := make(chan bool)
	wg.Add(2)
	go printOdd(oddChan, evenChan, &wg)
	go printEven(oddChan, evenChan, &wg)
	evenChan <- true
	wg.Wait()
	fmt.Println("Done")
}
