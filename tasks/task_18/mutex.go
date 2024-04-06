// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
}

func (c *counter) increase(m *sync.Mutex) {
	m.Lock()
	c.count++
	m.Unlock()
}

func main() {
	var m sync.Mutex
	var wg sync.WaitGroup

	c := &counter{count: 0}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.increase(&m)
		}()
	}

	wg.Wait()

	fmt.Printf("final counter value: %d", c.count)

}
