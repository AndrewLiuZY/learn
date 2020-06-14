package main

import "fmt"

import "strconv"

func main() {
	ch := make(chan int)
	end := make(chan int)
	go func() {
		for index := 1; index < 10; index++ {
			ch <- index * 10
		}
		end <- 0
	}()
	go func() {
		for index := 1; index < 10; index++ {
			fmt.Println(strconv.Itoa(<-ch))
		}
	}()
	<-end
}
