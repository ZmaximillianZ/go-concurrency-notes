package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}

	close(c) // close channel
}

func main() {
	fmt.Println("Start")
	c := make(chan int, 19)

	go squares(c) // start goroutine

	// periodic block/unblock of main goroutine until chanel closes
	// Go предоставляет ключевое слово range, которое автоматически
	// останавливает цикл, когда канал будет закрыт.
	for val := range c {
		fmt.Println(val)
	}

	fmt.Println("End")
}
