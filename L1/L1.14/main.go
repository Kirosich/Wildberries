package main

import (
	"fmt"
)

func main() {
	var Value interface{}
	//Value = 10
	//Value = "teq"
	//Value = true
	Value = make(chan int)
	switch Value.(type) {
	case int:
		fmt.Println("Int type")
	case string:
		fmt.Println("String type")
	case bool:
		fmt.Println("Bool type")
	case chan int:
		fmt.Println("Chan int type")
	}
}
