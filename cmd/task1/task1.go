package main

import "fmt"

type Human struct {
	Name       string
	Age        int
	Occupation string
}

func (h *Human) Introducing() {
	fmt.Printf("Hello! My name is %v, I'm %v old and my occupation is %v.\n", h.Name, h.Age, h.Occupation)
}

func (h *Human) Addition(num1 int, num2 int) int{
	return num1 + num2
}


// Встраивание методов структуры Human в структуру Action
type Action struct {
	Human
	Num1 int
	Num2 int
}

func main() {
	// Создание переменной типа структуры Action и использование методов встроенной в него структуры Human
	action := Action{
		Human: Human{
			Name:       "Foo Barovich",
			Age:        25,
			Occupation: "programmer",
		},
		Num1: 500,
		Num2: 1000,
	}
	action.Introducing()
	fmt.Println(action.Addition(action.Num1, action.Num2))
}
