//Реализовать пересечение двух неупорядоченных множеств.

// пересечение множеств - элементы которые есть в обоих множествах

package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 5, 5, 5, 29, 6, 0, 2, 7}
	arr2 := []int{5, 5, 56, 6, 7, 7, 8, 0, 9, 12, 13}

	arr1Map := make(map[int]int)
	arr2Map := make(map[int]int)
	var res []int
	var count int

	for _, val := range arr1 {
		_, ok := arr1Map[val]
		if !ok {
			arr1Map[val] = 1
		} else {
			arr1Map[val] += 1
		}
	}

	for _, val := range arr2 {
		_, ok := arr2Map[val]
		if !ok {
			arr2Map[val] = 1
		} else {
			arr2Map[val] += 1
		}
	}

	for val := range arr1Map {
		_, ok := arr2Map[val]
		if ok {
			if arr1Map[val] > arr2Map[val] {
				count = arr2Map[val]
			} else {
				count = arr1Map[val]
			}
			for i := 0; i < count; i++ {
				res = append(res, val)
			}

		}
	}

	fmt.Println(res)

}
