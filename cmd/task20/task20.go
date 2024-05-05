package main

import (
	"fmt"
	"strings"
)

func main() {
	words := "snow dog sun"

	// Разделяем строку по пробелу и составляем слайс строк
	sliceString := strings.Split(words, " ")
	
	// Переворачиваем слайс
	for i := 0; i < len(sliceString)/2; i++ {
		sliceString[i], sliceString[len(sliceString)-1-i] = sliceString[len(sliceString)-1-i], sliceString[i]
	}

	// Соединяем с разделителем в виде пробела
	finalWords := strings.Join(sliceString, " ")

	fmt.Printf("%v - %v\n", words, finalWords)
}
