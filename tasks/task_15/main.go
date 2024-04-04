//К каким негативным последствиям может привести данный фрагмент кода,
//и как это исправить? Приведите корректный пример реализации.

// глобальные переменные в целом лучше не использовать тк
// мы не знаем что происходило со строкой в других частях пакета ->
// она может быть меньшей длины чем 100
// var justString string

// func someFunc() {
// 	// побитовый сдвиг влево = 1024
// 	v := createHugeString(1 << 10)
// 	//возможет выход за пределы макс индекса строки что приводит к панике
// 	justString = v[:2]
// }

// func main() {
// 	someFunc()
// }

package main

import (
	"fmt"
	"strconv"
)

func createHugeString(n int) string {
	return strconv.Itoa(n)
}

func someFunc(str string) string {
	v := createHugeString(1 << 10)
	//проверка на выход за границы строки
	if len(v) < 100 {
		str = v
	} else {
		str = v[:100]
	}
	return str
}

func main() {
	// лучше передать строку в качестве аргумента
	// вместо объявления ее глобальной переменной
	var justString string
	justString = someFunc(justString)

	fmt.Println(justString)
}
