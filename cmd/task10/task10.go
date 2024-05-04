package main

import "fmt"

func main() {
	temperatures := [8]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// Храним результат в мапе из слайсов
	steppedTemp := make(map[int][]float32)

	for _, temperature := range temperatures {
		// Вычтем первый разряд числа, чтобы найти шаг
		step := int(temperature) - int(temperature)%10
		steppedTemp[step] = append(steppedTemp[step], temperature)
	}

	fmt.Println(steppedTemp)
}
