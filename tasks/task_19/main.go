//Разработать программу, которая переворачивает подаваемую на ход строку
//(например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"fmt"
)

func reverse(str string) string {
	conv := []rune(str)

	for i, j := 0, len(conv)-1; i < j; i, j = i+1, j-1 {
		conv[i], conv[j] = conv[j], conv[i]
	}
	return string(conv)
}

func main() {
	str := "главрыба"
	str = reverse(str)
	fmt.Println(str)
}
