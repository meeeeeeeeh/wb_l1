//Разработать программу, которая проверяет, что
//все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// 	aabcd — false

package main

import (
	"fmt"
)

func uniq(str string) bool {
	strCheck := make(map[rune]rune)
	res := true

	for _, val := range str {
		_, ok := strCheck[val]
		if !ok {
			strCheck[val] = val
		} else {
			res = false
			break
		}
	}

	return res
}

func main() {
	str1 := "aAbojoihi"
	str2 := "bhgwfy"
	str3 := "hhopkm"

	fmt.Println(uniq(str1)) //false
	fmt.Println(uniq(str2)) //true
	fmt.Println(uniq(str3)) //false
}
