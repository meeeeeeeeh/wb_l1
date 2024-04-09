// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

//Объединяя несколько каналов и горутин, вы получаете возможность создавать потоки данных,
//которые в терминологии Go также называются конвейерами.

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func write(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
}

func read(ch <-chan int) {
	for i := range ch {
		fmt.Println(i) // <- ch - вычитывает значение из канала = i
		//<-ch
	}
}

func main() {
	secAmount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	ch := make(chan int)

	go write(ch)
	go read(ch) // можно го не писать тк дальше нет функций - после го не ожидается завершение функции и программа идет дальшн

	select {
	case <-time.After(time.Duration(secAmount) * time.Second):
		close(ch)
	}
}
