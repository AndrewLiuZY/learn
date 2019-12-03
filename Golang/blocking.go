package main

import (
	"fmt"
	"time"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int, 1)
	out <- 2
	go f1(out)
	time.Sleep(1e9)
}
