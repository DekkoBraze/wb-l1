package main

import "fmt"

func binarySearch(arr []int, searchedValue int) int {
	leftBorder := 0
	rightBorder := len(arr) - 1
	for leftBorder <= rightBorder {
		midBorder := (leftBorder + rightBorder) / 2
		midValue := arr[midBorder]

		if midValue == searchedValue {
			return midBorder
		} else if midValue < searchedValue {
			leftBorder = midBorder + 1
		} else {
			rightBorder = midBorder - 1
		}
	}

	return -1
}

func main() {
	exampleArr := []int{1, 5, 7, 10, 15, 20, 35}
	var exampleValue int
	fmt.Println("Введите искомое значение:")
	fmt.Scanln(&exampleValue)
	searchedIndex := binarySearch(exampleArr, exampleValue)
	if searchedIndex == -1 {
		fmt.Println("Значение не найдено.")
	} else {
		fmt.Printf("Индекс значения %v в списке - %v.\n", exampleValue, searchedIndex)
	}
}
