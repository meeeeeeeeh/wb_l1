// Реализовать паттерн «адаптер» на любом примере.

//паттерн позволяет разным объектам использовать один функционал
package main

import (
	"fmt"
)

type workerAdapter interface {
	action()
}

type counter struct{}

func (c *counter) count(n int) {
	fmt.Println("counts to", n)
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
} 

type writer struct{}

func (w *writer) write() {
	fmt.Println("writes smth")
}

type counterAdapter struct{
	*counter
}

type writerAdapter struct {
	*writer
}

func (adapter *counterAdapter) action() {
	adapter.count(5)
}

func newCounterAdapter(counter *counter) workerAdapter {
	return &counterAdapter{counter}
}

func (adapter *writerAdapter) action() {
	adapter.write()
}

func newWriterAdapter(writer *writer) workerAdapter {
	return &writerAdapter{writer}
}



func main() {
	w := newWriterAdapter(&writer{})
	c := newCounterAdapter(&counter{})

	w.action()
	c.action()
}
