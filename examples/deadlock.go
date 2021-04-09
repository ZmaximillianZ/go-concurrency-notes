package main

import "fmt"

func main() {
	fmt.Println("main() started")
	c := make(chan string)

	// горутина для чтения из канала
	//go func(ch chan string) {
	//	fmt.Println("read from chan ", <-ch)
	//	close(ch)
	//}(c)

	c <- "John"
	fmt.Println("main() stopped")
}

// output:
//main() started
//fatal error: all goroutines are asleep - deadlock!
//goroutine 1 [chan send]:
