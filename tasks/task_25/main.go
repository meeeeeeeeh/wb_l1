//Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func sleep(sec time.Duration) {
	for {
		select {
		case <-time.After(sec * time.Second):
			return
		}
	}
}

func main() {
	var t time.Duration
	t = 2

	fmt.Printf("waitting for %d seconds since %v\n", t, time.Now())

	sleep(t)
	fmt.Printf("finish waitting at %v)", time.Now())

}
