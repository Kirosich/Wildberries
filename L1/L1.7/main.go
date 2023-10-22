package main

import (
	"fmt"
	"sync"
	"time"
)

func Set(data *map[int]int, key int, value int, mu *sync.RWMutex) {
	mu.Lock()
	(*data)[key] = value
	mu.Unlock()
}
func Get(data *map[int]int, key int, mu *sync.RWMutex) {
	mu.RLock()
	value := (*data)[key]
	mu.RUnlock()
	fmt.Println("Key:", key, "Value:", value)
}

func main() {
	var mu sync.RWMutex
	data := make(map[int]int)
	key := 2
	value := 5
	go Set(&data, key, value, &mu)
	key++
	value++
	go Set(&data, key, value, &mu)
	time.Sleep(1 * time.Second)
	go Get(&data, key, &mu)
	key--
	go Get(&data, key, &mu)
	time.Sleep(1 * time.Second)
}
