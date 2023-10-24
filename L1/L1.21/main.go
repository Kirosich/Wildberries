package main

import "fmt"

type oldsum struct {
	num1 int
	num2 int
}

func (s oldsum) sum() int {
	return s.num1 + s.num2
}

type newsum interface {
	newsum(int, int)
}

type adapter struct {
	oldsum *oldsum
}

func (a *adapter) newsum(num1 int, num2 int) int {
	a.oldsum.num1 = num1
	a.oldsum.num2 = num2
	return a.oldsum.sum()
}

func main() {
	oldsum := &oldsum{
		num1: 10,
		num2: 12,
	}
	adapter := adapter{
		oldsum: oldsum,
	}
	a := adapter.newsum(adapter.oldsum.num1, adapter.oldsum.num2)
	fmt.Println(a)
}
