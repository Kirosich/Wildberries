package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human
}

func (h Human) Hello() {
	fmt.Println("Hello, I am", h.Name)
}

func main() {
	Human := Human{
		Name: "Ilya",
		Age:  20,
	}
	Action := &Action{
		Human: Human,
	}
	Action.Hello()
}
