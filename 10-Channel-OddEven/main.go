package main

import (
	"fmt"
)

func odd(ch chan int) {
	for i := 1; ; i += 2 {
		ch <- i
	}
}

func even(ch chan int) {
	for i := 2; ; i += 2 {
		ch <- i
	}
}

func main() {
	oddChan := make(chan int)
	evenChan := make(chan int)

	go odd(oddChan)
	go even(evenChan)

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(<-oddChan)
		} else {
			fmt.Println(<-evenChan)
		}
	}
}
