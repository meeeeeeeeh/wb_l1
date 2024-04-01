// Удалить i-ый элемент из слайса.
package main

import (
	"errors"
	"fmt"
	"log"
)

func remove(slice []int, idx int) ([]int, error) {
	if idx >= len(slice) || idx < 0 {
		err := errors.New("index out of range")
		return nil, err
	}

	res := make([]int, len(slice)-1)
	for i, j := 0, 0; i < len(slice); i, j = i+1, j+1 {
		if idx == i {
			i++
		}
		res[j] = slice[i]

	}

	return res, nil
}

func main() {
	s := []int{1, 2, 3, 4, 5}

	s, err := remove(s, 0)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(s)
}
