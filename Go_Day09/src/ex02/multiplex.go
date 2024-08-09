package main

import (
	"sync"
)

func multiplex(inputs ...<-chan interface{}) chan interface{} {
	result := make(chan interface{})
	wg := &sync.WaitGroup{}

	for _, input := range inputs {
		wg.Add(1)
		go func(input <-chan interface{}) {
			defer wg.Done()
			for data := range input {
				result <- data
			}
		}(input)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	return result

}
