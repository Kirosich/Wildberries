package main

import (
	"fmt"
	"sync"
)

func Square(num int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- num * num
}

func main() {
	var wg sync.WaitGroup
	var sum int
	nums := []int{2, 4, 6, 8, 10}
	numCh := make(chan int, len(nums))

	for _, val := range nums {

		wg.Add(1)
		go Square(val, numCh, &wg)
	}
	wg.Wait()
	close(numCh)
	for num := range numCh {
		sum += num
	}
	fmt.Println(sum)
}
