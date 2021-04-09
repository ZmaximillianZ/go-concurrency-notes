package main

import (
	"fmt"
	"time"
)

var start1 time.Time

func init() {
	start1 = time.Now()
}

func service3(c chan string) {
	//time.Sleep(2 * time.Second)
	c <- "Hello from service 1 1"
	c <- "Hello from service 1 2"
	c <- "Hello from service 1 3"
}

func service4(c chan string) {
	//time.Sleep(4 * time.Second)
	c <- "Hello from service 2 1"
	c <- "Hello from service 2 2"
	c <- "Hello from service 2 3"
}

func main() {
	fmt.Println("main() start1ed", time.Since(start1))

	chan1 := make(chan string, 3)
	chan2 := make(chan string, 3)

	go service3(chan1)
	go service4(chan2)

	// т.к. select не блокируется с default,
	// то для того чтобы планировщик переключил контекст на горутины
	// необходимо в main'е "заснуть"
	time.Sleep(1 * time.Millisecond)

LOOP:
	for {
		select {
		case res := <-chan1:
			fmt.Println("Response from service 1", res, time.Since(start1))
		case res := <-chan2:
			fmt.Println("Response from service 2", res, time.Since(start1))
		// select не блокируется с default
		default:
			break LOOP
		}
	}

	fmt.Println("main() stopped", time.Since(start1))
}
