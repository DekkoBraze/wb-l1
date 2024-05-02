package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}

	for _, number := range numbers {
		waitGroup.Add(1)
		go func(number int) {
			defer waitGroup.Done()
			fmt.Println(number * number)
		}(number)
	}

	waitGroup.Wait()
}
