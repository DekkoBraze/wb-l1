package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {
	// Ввод с консоли времени работы
	var seconds time.Duration
	fmt.Println("Укажите время работы: ")
	fmt.Scanln(&seconds)

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

	// Горутина, которая читает данные из канала и выводит их в stdout
	waitGroup.Add(1)
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

	// Завершение горутин если пришел сигнал о прерывании
	signal.Notify(os_signals, os.Interrupt)
	go func() {
		for sig := range os_signals {
			if sig.String() == "interrupt" {
				cancel()
			}
		}
	}()

	// Завершение горутин после истечения указанного времени
	time.Sleep(seconds * time.Second)
	cancel()
	waitGroup.Wait()
}
