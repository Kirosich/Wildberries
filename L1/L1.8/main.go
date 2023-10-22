package main

import (
	"fmt"
)

func main() {
	var num int64
	// Ввод исходного числа
	fmt.Scanln(&num)
	//Исходное число в битах
	fmt.Printf("%b\n", num)
	var i int
	// Ввод i-го бита
	fmt.Scanln(&i)
	// Преобразованное число в битах и десятичной
	num = num | (1 << i)
	fmt.Printf("%b\n", num)
	fmt.Println(num)
}
