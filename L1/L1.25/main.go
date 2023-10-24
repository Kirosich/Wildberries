package main

import (
	"fmt"
	"time"
)

func Sleep(seconds time.Duration) {
	<-time.After(seconds)
}

func main() {
	Sleep(3 * time.Second)
	fmt.Println("3 sec")
}
