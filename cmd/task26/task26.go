package main

import (
	"fmt"
	"strings"
)

func IsUnique(words string) bool{
	words = strings.ToLower(words)
	set := make(map[rune]struct{})
	for _, word := range words {
		set[word] = struct{}{}
	}
	return len(words) == len(set)
}

func main() {
	var words string
	fmt.Println("Введите строку:")
	fmt.Scanln(&words)

	fmt.Println(IsUnique(words))
}
