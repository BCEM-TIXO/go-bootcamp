package main

import (
	"fmt"
	"sync"
	"time"
)

func SleepSort(unsortedInt []int) (sortedInt chan int) {
	sortedInt = make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(unsortedInt))
	for _, v := range unsortedInt {
		go func(v int) {
			defer wg.Done()
			time.Sleep(time.Second * time.Duration(v))
			sortedInt <- v
		}(v)
	}
	go func() {
		wg.Wait()
		close(sortedInt)
	}()
	return sortedInt
}

func main() {
	ints := []int{3, 2, 1, 1}
	fmt.Println(ints)
	c := SleepSort(ints)
	for i := range c {
		fmt.Println(i)
	}
}
