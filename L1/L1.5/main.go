package main

import (
	"fmt"
	"sync"
	"time"
)

func Writer(mainCh *chan int, exitCh *chan int, wg *sync.WaitGroup) {
	var i int
	for {
		select {
		case *mainCh <- i:
			i++
		case <-*exitCh:
			wg.Done()
			return
		}
	}
}

func Reader(mainCh *chan int, exitCh *chan int, wg *sync.WaitGroup) {
	for {
		select {
		case <-*mainCh:
			fmt.Println(<-*mainCh)
		case <-*exitCh:
			wg.Done()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	var mainCh = make(chan int)
	var exitCh = make(chan int)
	var N time.Duration
	fmt.Println("Enter how much time program will be working")
	fmt.Scanln(&N)
	wg.Add(2)
	go Writer(&mainCh, &exitCh, &wg)
	go Reader(&mainCh, &exitCh, &wg)
	time.Sleep(N * time.Second)
	close(exitCh)
	wg.Wait()
	fmt.Printf("it's been %v seconds", int(N))
}
