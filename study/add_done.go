package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	count := 20
	fmt.Printf("Going to create %d goroutines.\n", count)

	flag := true
	if len(os.Args) == 1 {
		flag = false
	}

	var waitGroup sync.WaitGroup

	fmt.Printf("%#v\n", waitGroup)
	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	// More Add() calls - c аргументами строки
	if flag {
		waitGroup.Add(1)

		// fatal error: all goroutines are asleep - deadlock!. Это означает, что
		//программа обречена бесконечно ждать завершения горутины,
		//то есть вызова Done(), которого никогда не будет.

	} else {
		waitGroup.Done() // More Done() calls - без аргументов

		// panic: sync: negative WaitGroup counter вызвано
		//количеством вызовов Done(), превышающим количество вызовов Add()
	}

	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}
