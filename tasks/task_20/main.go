//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	var res string

	div := strings.Split(str, " ")

	for i := len(div) - 1; i >= 0; i-- {
		res += div[i]
		res += " "
	}

	fmt.Println((res))

}

// функции склеивания или оператора сложения для конкатенации строк — не самое лучшее решение.
//Такой род операций каждый раз создает новую строку, снижая производительность.

// Поэтому в Golang есть специальный оптимизированный инструмент для с
//бора целых строк из составных частей с указанием неких правил — Builder:

// builded := &strings.Builder{}

// builded.WriteString("very")
// builded.WriteString(" ")
// builded.WriteString("long")
// builded.WriteString(" ")
// builded.WriteString("line")

// fmt.Println(builded.String()) // ВЫВОД: very long line
