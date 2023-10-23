package main

import "fmt"

func main() {
	var a, b int = 2, 5
	b, a = a, b
	fmt.Println(a, b)
}
