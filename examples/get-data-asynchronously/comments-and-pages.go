package main

import (
	"fmt"
	"time"
)

func getComments() chan string {
	// надо использовать буферизированный канал, это дает
	// записать нужное количество значений в канал, не блокируясь
	result := make(chan string, 1)
	go func(out chan<- string) {
		time.Sleep(2 * time.Second)
		fmt.Println("async operation ready, return comments")
		out <- "32 comments"
	}(result)
	return result
}
func getPage() {
	// вызов канала комментариев
	resultCh := getComments()
	time.Sleep(1 * time.Second)
	fmt.Println("get related articles")
	// ожидание комментариев
	commentsData := <-resultCh
	fmt.Println("main goroutine:", commentsData)
}

func main() {
	getPage()
}
