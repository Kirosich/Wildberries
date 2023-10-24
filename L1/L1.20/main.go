package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	row, _ := reader.ReadString('\n')
	slice := strings.Split(strings.TrimSuffix(row, "\n"), " ")
	for i := len(slice) - 1; 0 <= i; i-- {
		fmt.Print(slice[i], " ")
	}
}
