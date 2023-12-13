package Flags

import (
	"GrepTask/helpers"
	"bufio"
	"os"
	"strings"
)

func OpenFile(flags *Flags) [][]string {
	var lines [][]string

	file, err := os.Open(flags.FileName)
	helpers.HandleError(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, strings.Fields(scanner.Text()))
	}

	return lines
}
