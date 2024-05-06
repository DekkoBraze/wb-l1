package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Индекс удаляемого элемента
	index := 5

	// Создание нового слайса без удаляемого элемента
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println(slice)
}
