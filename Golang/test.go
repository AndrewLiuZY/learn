package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	t := time.Now()
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			for c := 0; c < 100; c++ {
				i++
				if 3*a+3*b == 21 && 3*a+b+2*c == 19 && 2*a+2*b+c == 15 {
					fmt.Println("猪：", a)
					fmt.Println("帽子：", b)
					fmt.Println("车：", c)
					fmt.Println("猪+二帽子*车：", a+2*b*c)
					fmt.Println("cal times：", i)
					fmt.Println("time：", time.Since(t).Seconds())
					return
				}
			}
		}
	}
}
