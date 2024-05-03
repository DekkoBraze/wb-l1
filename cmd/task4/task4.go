package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	var num_workers int
	fmt.Println("Введите количество воркеров: ")
	fmt.Scanln(&num_workers)

	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan string)
	os_signals := make(chan os.Signal, 1)

	// Горутина, осуществляющая запись в канал
	waitGroup.Add(1)
	go func(ch chan<- string) {
		defer waitGroup.Done()
		for {
			ch <- "New message to channel!"
			if ctx.Err() != nil {
				close(ch)
				return
			}
		}
	}(ch)

	// Воркеры, которые читают данные из канала и выводят в stdout
	waitGroup.Add(num_workers)
	for i := 0; i < num_workers; i++ {
		go func(ch <-chan string) {
			defer waitGroup.Done()
			for {
				message, ok := <-ch
				if ok {
					fmt.Println(message)
				} else {
					return
				}
			}
		}(ch)
	}

	// Обработка сигнала о прекращении работы
	signal.Notify(os_signals, os.Interrupt)
	go func() {
		for sig := range os_signals {
			if sig.String() == "interrupt" {
				cancel()
			}
		}
	}()

	waitGroup.Wait()
}
