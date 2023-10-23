package main

import (
	"fmt"
	"sort"
)

func main() {
	var mas = []int{1, 2, 10, 30, 2, 5, 9}
	sort.Slice(mas, func(i, j int) bool {
		return mas[i] < mas[j]
	})
	fmt.Println(mas)
}
