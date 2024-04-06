//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"fmt"
)

func quicksort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	partition := arr[len(arr)/2]

	var lower, bigger, equal []int

	for _, val := range arr {
		if val > partition {
			bigger = append(bigger, val)
		} else if val < partition {
			lower = append(lower, val)
		} else {
			equal = append(equal, val)
		}
	}

	lower = quicksort(lower)
	bigger = quicksort(bigger)
	return append(append(lower, equal...), bigger...)

}

func main() {
	arr := []int{7, 2, 0, 4, 1, 9, 8, 8, 5}

	arr = quicksort(arr)

	fmt.Println(arr)

}
