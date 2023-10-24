package main

import (
	"fmt"
	"strings"
)

func main() {
	var row string
	fmt.Scanln(&row)
	unic := make(map[string]string)
	for _, val := range row {
		letter := string(val)
		letter = strings.ToLower(letter)
		if letter == unic[letter] {
			fmt.Println(false)
			return
		}
		unic[letter] = letter
	}
	fmt.Println(true)
}
