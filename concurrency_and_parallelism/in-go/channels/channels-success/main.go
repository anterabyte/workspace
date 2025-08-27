package main

import (
	"fmt"
)

func main() {

	//c := make(chan int)

	// A succesful channel as it takes input from one side and ouput from other side at the same time.
	//go func() {
	//	c <- 42
	//}()

	//fmt.Println(<-c)

	// Simple even and odd channel

	even := make(chan int)

	go func() {

		for i := 0; i < 10; i++ {

			even <- i

		}

		close(even)

	}()

	for i := range even {
		if i%2 != 0 {
			fmt.Println("odd:", i)
		} else {
			fmt.Println("even:", i)
		}
	}
}
