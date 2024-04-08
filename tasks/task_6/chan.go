//Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"fmt"
)

func write() {
	for {
		fmt.Println("a")
	}

}

// func stop(ch chan bool) {

// }

func main() {
	ch := make(chan int)
	stop := make(chan bool)

	fmt.Println("Doing smth")
	go write()

	for i := 0; i < 10; i++ {

		select {
		//case ch <- i:

		case <-stop:
			close(ch)
			return
		}
	}

	stop <- true
	fmt.Println("Stop")

}
