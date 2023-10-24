package main

import "fmt"

func main() {
	var words []string
	words = []string{"worm", "work", "wok"}
	i := 1
	words = append(words[:i], words[i+1:]...)
	fmt.Println(words)
}
