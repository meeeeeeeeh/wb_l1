//Разработать конвейер чисел. Даны два канала:
//в первый пишутся числа (x) из массива,
//во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
	"sync"
)

func writeNum(num chan<- int, val int) {
	num <- val
	fmt.Println("writting: ", val)
}

func writeSq(num <-chan int, sq chan<- int) {
	x := <-num
	sq <- x * x
}

func printSq(sq <-chan int) {
	fmt.Println(<-sq)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	num := make(chan int)
	square := make(chan int)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(arr))

	// если функции будут не в горутинах то
	// анонимная функция запустится несколько раз паралллельно
	// 1 раз отработает но на втором канал будет заблокирован на запись
	//(тк небуф канал блокируется после записи значения до чтения)
	// -> deadlock
	for _, val := range arr {
		go func(val int) {
			go writeNum(num, val)
			go writeSq(num, square)
			go printSq(square)
			defer waitGroup.Done()
		}(val)
	}

	waitGroup.Wait()
	fmt.Println("done")
}
