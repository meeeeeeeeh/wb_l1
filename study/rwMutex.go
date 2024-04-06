package main

import (
	"fmt"
	"sync"
	"time"
)

// может быть несколько считывателей, владеющих мьютексом sync.RWMutex
// пока все считыватели мьютекса sync.RWMutex не разблокируют его,
// вы не сможете заблокировать его для записи
var Password *secret
var wg sync.WaitGroup

type secret struct {
	RWM      sync.RWMutex
	password string
}

func Change(pass string) {
	fmt.Println("Change() function")
	//Lock() и Unlock(), используемые в sync.Mutex когда нужно заблокировать и разблокировать мьютекс для записи.
	Password.RWM.Lock()
	fmt.Println("Change() Locked")
	time.Sleep(4 * time.Second)
	Password.password = pass
	Password.RWM.Unlock()
	fmt.Println("Change() UnLocked")
}

func show() {
	defer wg.Done()
	//RLock() и RUnlock() - используются для блокировки и разблокировки мьютекса для чтения.
	//нельзя вносить изменения в какие-либо общие переменные внутри блоков кода RLock() и RUnlock()
	Password.RWM.RLock()
	fmt.Println("Show function locked!")
	time.Sleep(2 * time.Second)
	fmt.Println("Pass value:", Password.password)
	defer Password.RWM.RUnlock()
}

func main() {
	Password = &secret{password: "myPass"}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go show()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		Change("123456")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		Change("54321")
	}()

	wg.Wait()

	// Direct access to Password.password
	fmt.Println("Current password value:", Password.password)
}
