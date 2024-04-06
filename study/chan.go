package main

import (
	"fmt"
	"sync"
)

// Эта функция просто записывает значение в канал и немедленно его закрывает
func writeToChannel(c chan int, x int) {
	c <- x
	close(c)
}

// Эта функция отправляет значение true в канал с типом bool.
func printer(ch chan bool) {
	ch <- true
}

func main() {
	// буферизированный канал. Как только буфер заполнится
	// мы сможем закрыть канал и при это горутина продолжит выполнение и вернется
	// небуферизированный канал блокируется пока значение из нео не будет кем то получено
	c := make(chan int, 1)

	//Здесь мы считываем из канала и выводим значение, не сохраняя его в отдельной переменной.
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func(c chan int) {
		defer waitGroup.Done()
		writeToChannel(c, 10)
		fmt.Println("Exit.")
	}(c)

	fmt.Println("Read:", <-c)

	//Код выше показывает метод определения того, закрыт канал или нет.
	//В этом случае мы игнорируем читаемое значение — если бы канал был открыт,
	//то значение было бы отброшено.
	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}

	waitGroup.Wait()

	//Здесь мы создадим небуферизованный канал и пять горутин без какой-либо
	//синхронизации, поскольку не будем использовать вызовы Add().

	// канал не закроется сразу же после записи????
	var ch chan bool = make(chan bool)
	for i := 0; i < 5; i++ {
		go printer(ch)
	}

	// диапазон по каналам
	// ВАЖНО: поскольку канал c не закрыт‚
	// цикл по диапазону не завершается сам по себе
	n := 0
	for i := range ch { // цикл завершится только если канал закрыт или использован брейк
		fmt.Println(i)
		if i == true {
			n++
		}
		if n > 2 {
			//Мы закрываем канал ch при выполнении условия и выходим из цикла for с по- мощью break.
			fmt.Println("n:", n)
			close(ch)
			break
		}
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch) // при чтение из закрытого канала получаем нулевое значение его типа данных - false
		// при записи в закрытый канал возникает паника

		// 4 аксиомы каналов!!
	}
}

// в какой то момент приводит к сбою:
//(глава про состояние гонки)
// panic: send on closed channel

// goroutine 8 [running]:
// main.printer(0x0?)
//         /Users/endeharh/wb_l1/study/chan.go:16 +0x1a
// created by main.main in goroutine 1
//         /Users/endeharh/wb_l1/study/chan.go:54 +0x24e
// exit status 2

//go run -race:
// ==================
// WARNING: DATA RACE
// Write at 0x00c00010e010 by main goroutine:
//   runtime.closechan()
//       /usr/local/go/src/runtime/chan.go:357 +0x0
//   main.main()
//       /Users/endeharh/wb_l1/study/chan.go:69 +0x4fd

// Previous read at 0x00c00010e010 by goroutine 10:
//   runtime.chansend()
//       /usr/local/go/src/runtime/chan.go:160 +0x0
//   main.printer()
//       /Users/endeharh/wb_l1/study/chan.go:16 +0x30
//   main.main.func3()
//       /Users/endeharh/wb_l1/study/chan.go:54 +0x33

// Goroutine 10 (running) created at:
//   main.main()
//       /Users/endeharh/wb_l1/study/chan.go:54 +0x32b
// ==================
// false
// false
// false
// false
// false
// panic: send on closed channel

//В строке 54 в channels.go происходит закрытие канала, а в строке 14 мы видим запись
//в тот же канал, которая, похоже, и является причиной проблемы с со- стоянием гонки.
//Строка 54 содержит close(ch), тогда как строка 14 — ch <- true.

//Состояние гонки данных — это ситуация, когда два или более запущенных элемента,
//таких как потоки и горутины, пытаются взять
//под контроль или изменить общий ресурс или общую переменную программы.
