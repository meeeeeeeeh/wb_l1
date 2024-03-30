package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	arr := []int{2, 4, 6, 8, 10}
	res := make([]int, len(arr))

	var waitGroup sync.WaitGroup
	for i, _ := range arr {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			res[i] = arr[i] * arr[i]
		}(i)
	}

	waitGroup.Wait()
	fmt.Println(res)
}
