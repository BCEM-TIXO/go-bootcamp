package main

import (
	"fmt"
)

const TotalGoroutines = 15

func main() {
	input1 := make(chan interface{})
	input2 := make(chan interface{})
	input3 := make(chan interface{})

	go func() {
		for i := 0; i < TotalGoroutines/3; i++ {
			input1 <- i
			input2 <- i * 10
			input3 <- i * 100
		}
		close(input1)
		close(input2)
		close(input3)
	}()

	fanIn := multiplex(input1, input2, input3)

	c := 0
	for data := range fanIn {
		c++
		fmt.Println(data)
	}

	if c == TotalGoroutines {
		fmt.Printf("--- All goroutines (%d) are proceeded ---\n", c)
	} else {
		fmt.Printf("--- Unexpected error, (%d) goroutines are proceeded ---\n", c)
	}
}
