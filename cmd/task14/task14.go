package main

import (
	"fmt"
	"reflect"
)

func main() {
	guesser(10)
	guesser("string")
	guesser(true)
	ch := make(chan int)
	guesser(ch)
}

func guesser(irfc interface{}) {
	fmt.Printf("%v - это %v!\n", irfc, reflect.TypeOf(irfc))
}

// Дополнительно: для решения задачи можно использовать специальную конструкцию со switch
func guesserWithSwitch(irfc interface{}) {
	switch v := irfc.(type) {
	case int:
		fmt.Printf("%v - это число!\n", v)
	case string:
		fmt.Printf("%v - это строка!\n", v)
	case bool:
		fmt.Printf("%v - это булевый тип!\n", v)
	case chan interface{}:
		fmt.Printf("%v - это интерфейсный канал!\n", v)
	}
}
