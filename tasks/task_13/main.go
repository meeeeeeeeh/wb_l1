//Поменять местами два числа без создания временной переменной.

package main

import (
	"fmt"
)

func main() {
	a := 2
	b := 5

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("a: %d, b: %d", a, b)
}
