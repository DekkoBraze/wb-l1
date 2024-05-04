package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	initNumbersChan := make(chan int)
	doubledNumbersChan := make(chan int)
	numbers := [5]int{1, 2, 3, 4, 5}

	// Горутина, посылающая числа в канал
	waitGroup.Add(1)
	go func(ch chan<- int) {
		defer waitGroup.Done()
		for _, number := range numbers {
			ch <- number
		}
		close(ch)
	}(initNumbersChan)

	// Горутина, производящая умножение и посылающая результат в канал
	waitGroup.Add(1)
	go func(numbersChan <-chan int, resultChan chan<- int) {
		defer waitGroup.Done()
		for {
			number, ok := <-numbersChan
			if ok {
				resultChan <- number * 2
			} else {
				close(resultChan)
				return
			}
		}
	}(initNumbersChan, doubledNumbersChan)

	// Горутина, посылающая данные из канала с результатами в stdout
	waitGroup.Add(1)
	go func(resultChan <-chan int) {
		defer waitGroup.Done()
		for {
			number, ok := <-resultChan
			if ok {
				fmt.Println(number)
			} else {
				return
			}
		}
	}(doubledNumbersChan)

	waitGroup.Wait()
}
