package main

import "fmt"

func great(c chan string) {
	fmt.Println(<-c) // читаем данные из канала, не блокируемая операция
	fmt.Println(<-c) // блокируемая операция т.к. пусто, планировщик переключается в main
}

func main() {
	fmt.Println("Start")
	// буферизированный канал
	ch := make(chan string, 1)
	// создаем горутину
	go great(ch)
	// блокирующая операция, т.к. буфер канала заполнен и планировщик переключается на
	// горутину great в которой будет производится операция чтения из канала
	ch <- "Max"
	// канал закрывается
	close(ch)
	// операция невозможна, т.к. канал закрыт
	// panic: send on closed channel
	ch <- "Den"
	fmt.Println("End")
}
