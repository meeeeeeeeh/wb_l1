//Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"errors"
	"fmt"
	"log"
)

// в функцию бинарного поиска должен подаваться уже отсортированный массив
func binarySearch(arr []int, num int) (int, error) {
	for len(arr) != 1 {
		index := len(arr) / 2
		if num == arr[index] {
			return arr[index], nil
		}
		if num > arr[index] {
			arr = arr[index:] // записывается все после указанного индекса
		} else {
			arr = arr[:index] // записывается все перед указанным индексом
		}
	}

	err := errors.New("No such value")
	return 0, err
}

func main() {
	numToFind := 2
	arr := []int{1, 2, 4, 5, 6, 7, 8, 9, 10}

	i, err := binarySearch(arr, numToFind)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(i)

}
