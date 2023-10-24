package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var row string
	reader := bufio.NewReader(os.Stdin)
	row, _ = reader.ReadString('\n')
	var slice []rune
	for _, val := range row {
		slice = append(slice, val)
	}
	for i := len(slice) - 1; 0 <= i; i-- {
		fmt.Printf("%c", slice[i])
	}
}
