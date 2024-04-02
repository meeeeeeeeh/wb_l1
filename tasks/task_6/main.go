//Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"fmt"
)

// func write() {
// 	fmt.Println("a")
// }

func main() {
	ch := make(chan int)
	stop := make(chan bool)

	for i := 0; ; i++ {
		select {
		case ch <- i:
		case <-stop:
			close(ch)
			return
		}
	}

	stop <- true

}
