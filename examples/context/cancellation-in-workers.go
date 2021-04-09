package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// запуск 10-ти воркеров и реализация отмены всех,
// кроме того который закончил работу первым

func worker(ctx context.Context, workerNum int, out chan<- int) {
	waitTime := time.Duration(rand.Intn(100)+10) * time.Millisecond
	fmt.Println(workerNum, "sleep", waitTime)
	select {
	// безопасный для использования между несколькими горутинами
	case <-ctx.Done():
		return
	case <-time.After(waitTime):
		fmt.Println("worker", workerNum, "done")
		out <- workerNum
	}
}

func main() {
	ctx, finish := context.WithCancel(context.Background())
	result := make(chan int, 1)
	for i := 0; i <= 10; i++ {
		go worker(ctx, i, result)
	}
	foundBy := <-result
	fmt.Println("result found by", foundBy)
	finish()
	time.Sleep(time.Second)
}
