package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * 1e9)
		ch <- 2
	}()
	<-ch
	fmt.Println("Done")

}
