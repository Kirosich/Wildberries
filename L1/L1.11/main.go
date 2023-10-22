package main

import "fmt"

func main() {
	mnoz := []int{5, 2, 3}
	mnoz2 := []int{2, 7, 5, 6}
	var peres []int
	for i := 0; i < len(mnoz); i++ {
		for j := 0; j < len(mnoz2); j++ {
			if mnoz[i] == mnoz2[j] {
				peres = append(peres, mnoz[i])
			}
		}
	}
	fmt.Println(peres)
}
