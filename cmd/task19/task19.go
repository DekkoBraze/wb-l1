package main

import "fmt"

func main() {
	var words string
	fmt.Println("Введите строку:")
	fmt.Scanln(&words)

	runesCount := 0
	runes := make([]rune, len(words))
	// Добавляем руны в слайс (в данном случае for идет по символам, а не по байтам)
	for _, symbol := range words {
		runes[runesCount] = symbol
		runesCount++
	}
	// Убираем нулевые руны
	runes = runes[0:runesCount]
	// Разворачиваем
	for i := 0; i < runesCount/2; i++ {
		runes[i], runes[runesCount-1-i] = runes[runesCount-1-i], runes[i]
	}
	// Преобразуем обратно в строку
	output := string(runes)
	fmt.Println(output)
}
