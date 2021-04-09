package main

import (
	"fmt"
	"time"
)

// Имитация веб сервиса в котором балансировщик нагрузки получает
// миллионы запросов и должен возвращать ответ от одной из доступных
// служб. Т.о мы запрашиваем у нескольких сервисов и тот кто
// ответил первый и будет использован

var start time.Time

func init() {
	start = time.Now()
}

// Так как service1 разблокируется раньше, чем service2, первый case
// разблокируется раньше и произведет чтение из chan1, а второй case
// будет проигнорирован.
func service1(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "Hello from service 1 1"
	c <- "Hello from service 1 2"
	c <- "Hello from service 1 3"
}

func service2(c chan string) {
	time.Sleep(4 * time.Second)
	c <- "Hello from service 2 1"
	c <- "Hello from service 2 2"
	c <- "Hello from service 2 3"
}

func main() {
	fmt.Println("main() started", time.Since(start))

	// чтобы сделать все блоки case неблокируемыми, мы можем использовать каналы с буфером.
	chan1 := make(chan string, 3)
	chan2 := make(chan string, 3)

	go service1(chan1)
	go service2(chan2)

	// Если все блоки case являются блокируемыми, тогда select будет ждать
	// до момента, пока один из блоков case разблокируется и будет выполнен.
	// Если несколько или все канальные операции не блокируемы, тогда один
	// из неблокируемых case будет выбран случайным образом (Примечание
	// переводчика: имеется ввиду случай, когда пришли одновременно данные
	// из двух и более каналов).
	select {
	// Так как каналы не используют буфер, операция чтения будет блокируемой.
	// Таким образом оба case будут блокируемыми и select будет ждать до тех
	// пор, пока один из case не разблокируется.
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
		// будет выполнятся default case
		//default:
		//	fmt.Println("Response Default case", time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
