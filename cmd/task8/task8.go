package main

import (
	"fmt"
	"strconv"
)

func main() {
	var index int
	var number int64
	fmt.Println("Введите индекс бита:")
	fmt.Scanln(&index)
	fmt.Println("Введите изменяемое число:")
	fmt.Scanln(&number)
	fmt.Println("До изменения:")
	fmt.Println(strconv.FormatInt(number, 2))
	// Сдвигаем бит на index и применяем операцию исключающего ИЛИ, чтобы сменить бит в числе на противоположный
	number ^= (1 << index)
	fmt.Println("После изменения:")
	fmt.Println(strconv.FormatInt(number, 2))
}
