package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
}

func (c *counter) inc(wg *sync.WaitGroup) {
	c.count++
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	counter := counter{count: 0}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go counter.inc(&wg)
	}
	wg.Wait()
	fmt.Println(counter.count)
}
