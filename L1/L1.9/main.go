package main

import "fmt"

func Send(mas []int, ch chan int) {
	for _, val := range mas {
		ch <- val
	}
	close(ch)
}

func Get(ch2 chan int, ch chan int) {
	for val := range ch {
		val = val * 2
		ch2 <- val
	}
	close(ch2)
}

func main() {
	mas := []int{1, 2, 3, 4}
	ch := make(chan int)
	ch2 := make(chan int)
	go Send(mas, ch)
	go Get(ch2, ch)
	for val := range ch2 {
		fmt.Println(val)
	}

}
