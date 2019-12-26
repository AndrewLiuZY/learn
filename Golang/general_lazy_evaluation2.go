package main

import (
	"fmt"
)

var resume chan int

func integers() chan int {
	yield := make(chan int)
	num1 := 0
	num2 := 1
	go func() {
		yield <- 0 //push first element
		for {
			num3 := num1 + num2
			yield <- num3
			num1 = num2
			num2 = num3
		}
	}()
	return yield
}

func generateInteger() int {
	return <-resume
}

func main() {
	resume = integers()
	for index := 0; index < 10; index++ {
		fmt.Println(generateInteger())
	}
}
