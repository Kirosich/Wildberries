package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	numbers := []int{2, 4, 6, 8, 10}

	for _, val := range numbers {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			fmt.Println(val * val)
		}(val)
	}
	wg.Wait()
}
