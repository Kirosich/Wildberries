package main

import "fmt"

func main() {
	mas := []string{"cat", "cat", "dog", "cat", "tree"}
	data := make(map[string]int)
	for _, val := range mas {
		if value, exists := data[val]; exists {
			data[val] = value + 1
		} else {
			data[val] = 1
		}
	}
	fmt.Println(data)
}
