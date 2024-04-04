//Разработать программу, которая в рантайме способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var n interface{}

	n = 1
	fmt.Println(reflect.TypeOf(n))

	n = "string"
	fmt.Println(reflect.TypeOf(n))

	n = true
	fmt.Println(reflect.TypeOf(n))

	n = make(chan int)
	fmt.Println(reflect.TypeOf(n))
}
