package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func gen(min, max int, createNumber chan int, end chan bool) {
	time.Sleep(time.Second)
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			fmt.Println("Ended!")
			return
		case <-time.After(4 * time.Second):
			//ожидает истечения указанного времени (4 * time.Second),
			//после чего выводит сообщение и правильно завершает gen() с помощью return
			fmt.Println("time.After()!")
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	createNumber := make(chan int)
	end := make(chan bool)

	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer!")
		return
	}

	n, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Going to create %d random numbers.\n", n)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		gen(0, 2*n, createNumber, end)
		wg.Done()
	}()

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", <-createNumber)
	}

	end <- true
	wg.Wait()
	fmt.Println("Exiting...")
}

//select позволяет прослушивать несколько каналов одновременно.
//Оператор select не вычисляется последовательно,
//так как все его каналы про- веряются одновременно.
//Если ни один из каналов в операторе не готов, то он блокируется (ожидает), пока один из каналов не будет готов
//Если несколько каналов оператора select готовы, то среда выполнения Go делает случайный выбор из набора этих готовых каналов.
//select без каких-либо разделов (select{}) запустит вечное ожидание.
