//Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	data := make(map[int]string)
	keyAmount := 5
	var val string
	var waitGroup sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < keyAmount; i++ {
		val = "str_" + strconv.Itoa(i)
		waitGroup.Add(1)

		go func(key int, val string) {
			defer waitGroup.Done()
			m.Lock()
			data[key] = val
			m.Unlock()
		}(i, val)

		//go writeData(data, i)
	}

	waitGroup.Wait()

	for k, v := range data {
		fmt.Println(k, v)
	}
}

// Мьютекс работает как буферизованный канал с пропускной способностью в единицу
//что позволяет не более чем одной горутине получать доступ к общей переменной в один момент времени
