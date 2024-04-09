//Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"fmt"
	"time"
)

func worker(stop chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("stop working")
			close(stop)
			return
		default:
			fmt.Println("working")
			time.Sleep(time.Second * 1)
		}
	}

}

func main() {
	stop := make(chan bool)

	go worker(stop)
	time.Sleep(time.Second * 5)

	stop <- true
	fmt.Println("program has finished")

}
