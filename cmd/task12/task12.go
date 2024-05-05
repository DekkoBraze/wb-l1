package main

import "fmt"

func main() {
	arr := [5]string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})
	properSet := make([]string, 0)

	for _, item := range arr {
		set[item] = struct{}{}
	}

	for item := range set {
		properSet = append(properSet, item)
	}

	fmt.Println(properSet)
}
