package main

import (
	"fmt"
)

// deadlock
func main() {

	ch := make(chan int)

	ch <- 1
	// ch <- 2
	// ch <- 3

	//fmt.Println(<-ch)
	//close(ch)

	// for i := 0; i < 2; i++ {
	// 	fmt.Printf("val: %d\n", <-ch)
	// }

	for val := range ch {
		fmt.Printf("val: %d\n", val)
	}

	//close(ch)

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

}
