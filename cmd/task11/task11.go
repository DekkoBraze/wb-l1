package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{5, 8, 1, 9, 3}
	count := make(map[int]int)
	intersection := make([]int, 0)
	mutex := sync.RWMutex{}

	// Горутина, которая идет по первому множеству 
	waitGroup.Add(1)
	go func(count map[int]int) {
		defer waitGroup.Done()
		for _, item := range arr1 {
			mutex.Lock()
			count[item] += 1
			mutex.Unlock()
		}
	}(count)
	
	// Горутина, которая идет по второму множеству 
	waitGroup.Add(1)
	go func(count map[int]int) {
		defer waitGroup.Done()
		for _, item := range arr2 {
			mutex.Lock()
			count[item] += 1
			mutex.Unlock()
		}
	}(count)

	waitGroup.Wait()

	// Выводим те элементы, которых насчитали ровно 2
	for number, item := range count {
		if item == 2 {
			intersection = append(intersection, number)
		}
	}

	fmt.Println(intersection)
}