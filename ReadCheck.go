package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan int)
	go func() {
		a <- 3
		close(a)
	}()
	go func() {
		for {
			v := <-a
			fmt.Println(v)
		}
	}()
	time.Sleep(time.Second)
}
