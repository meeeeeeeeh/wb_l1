package main

import (
	"context"
	"fmt"
	"time"
)

func print(cont context.Context) {
	for {
		select{
		case <-cont.Done():
			fmt.Println("stop working")
			return
		default :
			fmt.Println("working")
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {
	cont, cancel := context.WithCancel(context.Background())
	
	go print(cont)
	time.Sleep(time.Second * 5)

	cancel() //остановка горутины

	fmt.Println("program has finished")
}
