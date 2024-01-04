package channels

import (
	"fmt"
)

func Test() {

	channel := make(chan int)
	//? Pode ter um valor na qual o channel vai esperar encher para começar a ser lido
	// channel := make(chan int, 100)

	go set(channel)
	print(channel)

}

func set(channel chan<- int) {
	for i := 0; i < 100; i++ {
		channel <- i
		fmt.Println("Enviando: ", i)
	}
	close(channel)
}

func print(channel <-chan int) {
	for value := range channel {
		fmt.Println("Recebendo: ", value)
	}
}
