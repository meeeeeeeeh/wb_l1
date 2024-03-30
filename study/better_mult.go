package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	count := 10
	arguments := os.Args
	if len(arguments) == 2 {
		t, err := strconv.Atoi(arguments[1])
		if err == nil {
			count = t
		}
	}

	fmt.Printf("Going to create %d goroutines.\n", count)

	var waitGroup sync.WaitGroup //создается переменная sync.WaitGroup
	fmt.Printf("%#v\n", waitGroup)
	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		// Мы вызываем Add(1) непосредственно перед созданием горутины, чтобы избежать состояний гонки.
		// увеличивет счетчик на 1
		go func(x int) {
			defer waitGroup.Done()
			//Вызов Done() будет выполнен непосредственно перед возвратом анонимной функции
			// уменьшает счетчик на 1 -> делает Add(-1)
			fmt.Printf("%d ", x)
		}(i)
	}
	// планировщик определяет порядок запуска горутин

	// вайтгруп не ноль ли счетчик. После этого может закрываться мейн.
	//Иначе мейн закроется раньше чем отрабоотали горутины

	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()
	//Функция Wait() будет ждать, пока счетчик в переменной WaitGroup не станет
	//равным 0, после чего произойдет возврат, чего мы‚ собственно‚ и добиваемся.
	fmt.Println("\nExiting...") //Когда функция Wait() возвращается, выполняется оператор fmt.Println().
	//Таким образом‚ нам больше не нужно вызывать time.Sleep()!
}
