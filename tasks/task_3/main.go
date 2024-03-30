package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….)
// с использованием конкурентных вычислений.

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var res int

	var waitGroup sync.WaitGroup
	for i, _ := range arr {
		waitGroup.Add(1)

		go func(res *int, i int) {
			defer waitGroup.Done()
			*res += arr[i] * arr[i]
		}(&res, i)
	}

	waitGroup.Wait()
	fmt.Println(res)

}
