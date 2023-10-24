package main

import (
	"fmt"
)

func main() {
	a := (1 << 20) + 10
	b := (1 << 20) + 5
	if a > 1<<20 && b > 1<<20 {
		fmt.Println("a * b =", a*b)
		fmt.Println("a / b =", float64(a)/float64(b))
		fmt.Println("a + b =", a+b)
		fmt.Println("a - b =", a-b)
	} else {
		fmt.Println("Значения a и b должны быть больше 2^20")
	}
}
