package main

import "fmt"

func main() {
	var x int
	fmt.Println("Введите x: ")
	fmt.Scanln(&x)
	var y int
	fmt.Println("Введите y: ")
	fmt.Scanln(&y)

	x += y
	y = x - y
	x = x - y

	fmt.Printf("x = %v, y = %v\n", x, y)
}
