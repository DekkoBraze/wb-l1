package main

import (
	"fmt"
	"sync"

	"golang.org/x/sync/syncmap"
)

var waitGroup sync.WaitGroup

func main() {
	mapWithSyncMap()
}

// Конкурентная запись в map с помощью мьютекса
func mapWithMutex() {
	numbers := [5]int{1, 2, 3, 4, 5}
	dict := make(map[int]int)
	dictMutex := sync.RWMutex{}

	for _, number := range numbers {
		waitGroup.Add(1)
		go func(number int) {
			defer waitGroup.Done()
			// Одна горутина блокирует запись, остальные ждут разблокирования
			dictMutex.Lock()
			dict[number] = number
			dictMutex.Unlock()
		}(number)
	}

	waitGroup.Wait()
	fmt.Println(dict)
}

// Конкурентная запись в map с помощью типа syncmap.Map
func mapWithSyncMap() {
	numbers := [5]int{1, 2, 3, 4, 5}
	dict := syncmap.Map{}

	for _, number := range numbers {
		waitGroup.Add(1)
		go func(number int) {
			defer waitGroup.Done()
			dict.Store(number, number)
		}(number)
	}

	waitGroup.Wait()
	for _, number := range numbers {
		mapNumber, _ := dict.Load(number)
		fmt.Println(mapNumber)
	}
}
