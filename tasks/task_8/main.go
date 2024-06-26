//Дана переменная int64.
//Разработать программу которая устанавливает i-й бит в 1 или 0

package main

import (
	"fmt"
)

// & - побитовое и
// 0 & 0 = 0
// 0 & 1 = 0
// 1 & 0 = 0
// 1 & 1 = 1

// >> сдвиг вправо

// сначала сдвигаем вправо число на индекс таким образом первый бит - тот который нужно узнать
// далее операция & 001 после которой все биты кроме первого будут равны 0 в любом случае
// а первый юудет 1 если изначально был 1 и 0 если изначально был 0
func getBit(num int64, idx int64) int64 {
	return 1 & (num >> idx)
}

// | - побитовое или
// 0 | 0 = 0
// все остальное = 1
// сдвигаем 00001 на idx(2) -> 00100
// после пбитового или бит на месте индекса в любом случае будет равен 1
func setBit(num *int64, idx int64) {
	*num |= (1 << idx)
	//return num

}

// ^ - исключающее или (в го используется вместо инверсии ~)
func clearBit(num *int64, idx int64) {
	*num &^= 1 << idx
}

func main() {
	var num int64
	num = 7
	fmt.Printf("%064b\n", num)

	//fmt.Println(getBit(num, 0))

	setBit(&num, 5)
	fmt.Printf("%064b\n", num)

	clearBit(&num, 1)
	fmt.Printf("%064b\n", num)

}
