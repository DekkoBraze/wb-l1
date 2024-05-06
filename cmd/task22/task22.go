package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewInt(1 << 21)
	y := big.NewInt(1 << 25)
	fmt.Printf("x = %v\n", x)
	fmt.Printf("y = %v\n", y)
	res := big.NewInt(-1)
	res.Add(x, y)
	fmt.Printf("x + y: %v\n", res)
	res.Sub(x, y)
	fmt.Printf("x - y: %v\n", res)
	res.Mul(x, y)
	fmt.Printf("x * y: %v\n", res)
	res.Quo(x, y)
	fmt.Printf("x / y: %v\n", res)
	res.Rem(x, y)
	fmt.Print("x % y: ", res, "\n")
}