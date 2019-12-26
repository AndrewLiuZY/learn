package main

import (
	"flag"
	"time"

	"fmt"
)

var ngoroutine = flag.Int("n", 100000, "how many goroutines")

func f(left, right chan int) { left <- 1 + <-right }

func main() {
	t := time.Now()
	flag.Parse()

	leftmost := make(chan int)

	var left, right chan int = nil, leftmost

	for i := 0; i < *ngoroutine; i++ {

		left, right = right, make(chan int)

		go f(left, right)

	}

	right <- 0

	// start the chaining 开始链接

	x := <-leftmost // wait for completion 等待完成

	fmt.Println(x)
	fmt.Println(time.Since(t).Seconds())

	// 结果： 100000 ， 大约 1,5s （我实际测试只用了不到200ms）

}
