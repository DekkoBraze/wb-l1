package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(length int) string {
	return "stringstringstringstringstringstringstringstringstringstringstringstringstringstringstringstringstringstring"
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
	fmt.Println(justString)
}

/*
Пример из задания:
var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

В данном примере мы создаем большую строку (неизменяемый слайс байт), на её основе создаём другой слайс и присваиваем его другой строке (глобальной переменной).
Так как строка - это слайс, при присваивании она передает указатель на изначальный массив другой строке, не копируя содержание.
У такого подхода есть следующие недостатки:
1. В программе мы используем 100 первых символов из строки в 1024 символа (они остаются в памяти, а не очищаются после завершения работы функции)
2. Некоторые символы, которые не умещаются в один байт (н-р, Unicode), могут быть обрезаны и отображаться в новой переменной некорректно
Для того, чтобы решить проблему 1, мы можем копировать содержимое слайса с помощью strings.clean().
Для того, чтобы решить проблему 3, необходимо привести изначальный слайс бит к слайсу рун, однако это стоит делать только при необходимости, т к []rune
больше, чем []byte.
*/
