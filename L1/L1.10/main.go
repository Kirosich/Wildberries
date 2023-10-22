package main

import (
	"fmt"
)

func main() {
	var mas []float32
	var response string
	// Enter the numbers
	for {
		fmt.Print("Enter new temperature?(y/n):")
		fmt.Scanln(&response)
		if response == "y" {
			fmt.Println("Enter new number:")
			var num float32
			fmt.Scanln(&num)
			mas = append(mas, num)
		}
		if response == "n" {
			break
		}
	}

	data := make(map[int][]float32)
	for i := 0; i < len(mas); i++ {
		var groupKey int
		groupKey = int(mas[i] / 10)
		if groupKey == 0 {
			data[0] = append(data[0], mas[i])
		} else {
			groupKey *= 10
			data[groupKey] = append(data[groupKey], mas[i])
		}
	}
	fmt.Println(data)

}
