package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mutex sync.RWMutex
}

var waitGroup sync.WaitGroup

func main() {
	var desiredValue int
	fmt.Println("Введите искомое значение:")
	fmt.Scanln(&desiredValue)
	var workersNum int
	fmt.Println("Введите количество горутин:")
	fmt.Scanln(&workersNum)

	counter := Counter{0, sync.RWMutex{}}

	for i := 0; i < workersNum; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			for {
				// Используем мьютекс для того, чтобы избежать инкрементирования при уже достигнутом значении
				counter.mutex.Lock()
				if (counter.count != desiredValue) {
					counter.count += 1
				} else {
					counter.mutex.Unlock()
					return
				}
				counter.mutex.Unlock()
			}
		}()
	}

	waitGroup.Wait()
	fmt.Println(counter.count)
}
