package main

import (
	"fmt"
	"time"
)

func printme(x int) {
	fmt.Println("*", x)
	return
}

func main() {
	go func(x int) { // Так запускается анонимная функция в качестве горутины
		fmt.Printf("%d ", x)
	}(10) // (10) в конце — это способ, с помощью которого передается параметр анонимной функции.

	go printme(15)

	time.Sleep(time.Second)
	fmt.Println("Exiting...")
}

// вывод иногда 10 * 15
// но иногда * 15 10
