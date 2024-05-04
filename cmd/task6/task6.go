package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {
	stopGoWithVar()
}

// Остановка горутины с помощью передачи сигнала через канал
func stopGoWithChan() {
	quit := make(chan bool)

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		for {
			select {
			case <-quit:
				return
			default:
				fmt.Println("New message!")
			}
		}
	}()

	time.Sleep(5 * time.Second)
	quit <- true
	waitGroup.Wait()
}

// Остановка горутины с помощью контекста
func stopGoWithCtx() {
	ctx, cancel := context.WithCancel(context.Background())

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		for {
			fmt.Println("New message!")
			if ctx.Err() != nil {
				return
			}
		}
	}()

	time.Sleep(5 * time.Second)
	cancel()
	waitGroup.Wait()
}

// Остановка горутины с помощью закрытия канала
func stopGoWithCloseChan() {
	closeChan := make(chan struct{})

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		for {
			select {
			case <-closeChan:
				return
			default:
				fmt.Println("New message!")
			}
		}
	}()

	time.Sleep(5 * time.Second)
	close(closeChan)
	waitGroup.Wait()
}

// Остановка горутины с помощью глобальной/передаваемой по адресу переменной
func stopGoWithVar() {
	stopGo := false

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		for {
			fmt.Println("New message!")
			if stopGo {
				return
			}
		}
	}()

	time.Sleep(5 * time.Second)
	stopGo = true
	waitGroup.Wait()
}
