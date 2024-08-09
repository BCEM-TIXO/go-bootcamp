package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"sync"
)

func getUrlBody(url string) (string, error) {
	r, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func runWork(ctx context.Context, urls <-chan string, result chan<- *string) {
	wg := &sync.WaitGroup{}
	pool := make(chan struct{}, 8)
	for url := range urls {
		select {
		case <-ctx.Done(): // Прилетел сигнал завершения, сворачиваемся
			break
		case pool <- struct{}{}: // Есть работа, работаем
			wg.Add(1)
			go func(url string) {
				defer func() {
					wg.Done()
					<-pool
				}()
				body, err := getUrlBody(url)
				if err != nil {
					log.Println("Error reading body of %s: %v\n", url, err)
				}
				result <- &body
			}(url)
		}
	}
	wg.Wait()

	close(result)
}

func crawlWeb(ctx context.Context, urls <-chan string) <-chan *string {
	result := make(chan *string)
	go runWork(ctx, urls, result)
	return result
}
