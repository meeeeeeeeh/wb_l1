package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	name string
	age  int
}

func (h *Human) talk(word string) {
	fmt.Printf("Says %s,%s", word, h.name)
}

type Action struct {
	Human
}

func main() {
	var h Human
	h.name = "Pasha"
	act := &Action{Human: h}
	act.talk("what????")
}
