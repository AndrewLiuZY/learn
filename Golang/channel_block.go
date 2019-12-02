package main

import "fmt"

import "time"

func main() {
	ch1 := make(chan int)
	go pump(ch1) // pump hangs
	go suck(ch1)
	time.Sleep(1 * 1e9)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		fmt.Printf("send %d in channel\n", i)
		ch <- i
	}
}
func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
