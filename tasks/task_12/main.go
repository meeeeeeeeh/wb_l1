//Имеется последовательность строк -
//(cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import (
	"fmt"
)

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	res := make(map[string]string)

	for _, val := range strings {
		res[val] = val
	}

	fmt.Println(res)
}
