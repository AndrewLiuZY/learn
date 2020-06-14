package main

import "fmt"

import "strconv"

import "time"

func main() {
	ch := make(chan int)
	go tel(ch)
	setCloseTimer(ch, 1e9)
	for i := range ch {
		fmt.Println(strconv.Itoa(i))
	}

}

func setCloseTimer(ch chan int, i time.Duration) {
	go func() {
		time.Sleep(i)
		close(ch)
	}()
}

func tel(ch chan int) {
	defer func() {
		if _, ok := recover().(error); ok {
			fmt.Println("channel is closed  ")
		}
	}()
	i := 1
	for {
		ch <- i
		i++
	}
}
