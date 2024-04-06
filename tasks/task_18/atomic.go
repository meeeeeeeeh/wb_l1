//Атомарная операция — это операция, которая выполняется за один шаг относительно других
//потоков или, в данном случае, других горутин.
//->атомарную операцию нельзя прервать в середине ее работы
// это позволяет избежать состояния гонки без использования мьютексов

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type atomicCounter struct {
	val int32
}

// спользовании атомарной переменной во избежание состояний гонки
// все операции чтения и записи атомарной переменной
// должны выполняться с помощью функций, предоставляемых пакетом atomic.
func (c *atomicCounter) increase() {
	atomic.AddInt32(&c.val, 1)
}

func main() {
	var wg sync.WaitGroup

	c := atomicCounter{val: 0}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.increase()
		}()
	}

	wg.Wait()
	fmt.Printf("final counter value: %d", c.val)
	//go run -race atomic.go не выдает ошибок
}
