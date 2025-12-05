package main

import (
	"fmt"
	"sync"
)

//-------------------without wg------------------------------------
// func ping(pingCh, pongCh chan string) {
// 	for {
// 		msg := <-pingCh
// 		fmt.Println(msg)
// 		time.Sleep(time.Second)
// 		pongCh <- "Pong"
// 	}
// }

// func pong(pingCh, pongCh chan string) {
// 	for {
// 		msg := <-pongCh
// 		fmt.Println(msg)
// 		time.Sleep(time.Second)
// 		pingCh <- "Ping"
// 	}
// }
// func main() {
// 	pingCh := make(chan string)
// 	pongCh := make(chan string)
// 	go ping(pingCh, pongCh)
// 	go pong(pingCh, pongCh)
// 	pingCh <- "Ping"
// 	select {}
// }
//-------------------with wg------------------------------------
func ping(pingCh, pongCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		msg := <-pingCh
		fmt.Println(msg)
		pongCh <- "Pong"
	}
}

func pong(pingCh, pongCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		msg := <-pongCh
		fmt.Println(msg)
		if i < 4 {
			pingCh <- "Ping"
		}
	}
}
func main() {
	var wg sync.WaitGroup
	pingCh := make(chan string)
	pongCh := make(chan string)
	wg.Add(2)
	go ping(pingCh, pongCh, &wg)
	go pong(pingCh, pongCh, &wg)
	pingCh <- "Ping"
	wg.Wait()
	fmt.Println("Game over")
}
