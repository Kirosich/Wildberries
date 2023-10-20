package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func RecordCreater(mainCh *chan int, doneCh *chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(*mainCh)
	i := 0
	for {
		select {
		case *mainCh <- i:
			i++

		case <-*doneCh:
			return
		}
	}
}

func createWorkers(amount int, mainCh *chan int, doneCh *chan struct{}, wg *sync.WaitGroup) {
	for i := 0; i < amount; i++ {
		go Worker(mainCh, doneCh, i, wg)
	}
}

func Worker(mainCh *chan int, doneCh *chan struct{}, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case reader, ok := <-*mainCh:
			if ok != true {
				fmt.Println("Got close signal")
				return
			}
			fmt.Println("Goroutine id:", index, "Read from chan:", reader)
			time.Sleep(500 * time.Millisecond)
		case <-*doneCh:
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	var mainCh = make(chan int)
	var doneCh = make(chan struct{})
	var amountWorkers int

	//Sys signal handler
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	wg.Add(1)
	go RecordCreater(&mainCh, &doneCh, &wg)
	fmt.Println("Enter amount of workers: ")
	fmt.Scanln(&amountWorkers)
	createWorkers(amountWorkers, &mainCh, &doneCh, &wg)
	wg.Add(amountWorkers)
	fmt.Println("Workers is created...")

	//Wait for signal
	<-sigCh
	close(doneCh)
	wg.Wait()
	fmt.Println("WaitGroup is closed")
	os.Exit(0)
}
