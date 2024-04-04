//Дана последовательность температурных колебаний:
//-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов.
//Последовательность в подмножноствах не важна.

//Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

package main

import (
	"fmt"
)

func main() {
	var arr = []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float32)
	var key int

	for _, val := range arr {
		key = int(val)
		key = key / 10 * 10

		_, ok := groups[key]
		if !ok {
			newGroup := make([]float32, 1)
			newGroup[0] = val
			groups[key] = newGroup
		} else {
			groups[key] = append(groups[key], val)
		}

	}

	for k, v := range groups {
		fmt.Println("k, v: ", k, v)
	}
}
