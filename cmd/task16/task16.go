package main

import "fmt"

func main() {
	exampleArr := []int{5, 3, 4, 1, 100, -10, -15}
	fmt.Println(quickSort(exampleArr))
}

func quickSort(arr []int) (outArr []int){
	less := make([]int, 0)
	equal := make([]int, 0)
	greater := make([]int, 0)
	if len(arr) > 1 {
		// Задаем опорный элемент
		pivot := arr[0]
		for _, x := range arr {
			// Распределяем элементы на меньшие, равные и большие опорного
			if x < pivot {
				less = append(less, x)
			} else if x == pivot {
				equal = append(equal, x)
			} else if x > pivot {
				greater = append(greater, x)
			}
		}
		// Рекурсивно проходимся еще раз по меньшим и большим элементам
		finalLess := quickSort(less)
		finalGreater := quickSort(greater)
		// Собираем финальный слайс
		outArr = append(finalLess, equal...)
		outArr = append(outArr, finalGreater...)
		return
	} else {
		return arr
	}
}