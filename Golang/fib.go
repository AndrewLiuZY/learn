package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{10, 20, 30, 40, 50, 100, 500, 1000, 10000, 50000, 1000000, 10000000}
	for _, item := range s {
		testFib(item)
		fmt.Println("-------------------------------------------")
	}
}

func testFib(n int) {
	fmt.Println("n =", n)
	fmt.Print("fib1: ")
	recordTime(func() {
		fib1(n)
	})
	fmt.Print("fibMemo: ")
	recordTime(func() {
		fibMemo(n)
	})
	fmt.Print("fib3: ")
	recordTime(func() {
		fib3(n)
	})
}

//递归(自上向下)
func fib1(n int) int {
	if n <= 0 {
		return n
	}
	if n < 3 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

//备忘录递归(自上向下)
func fibMemo(n int) int {
	if n <= 0 {
		return 0
	}
	s := make([]int, 0)
	for i := 0; i < n+1; i++ {
		s = append(s, -1)
	}
	return fib2(s, n)
}

//备忘录递归(自上向下)
func fib2(s []int, n int) int {
	if s[n] != -1 {
		return s[n]
	}
	if n <= 2 {
		s[n] = 1
	} else {
		s[n] = fib2(s, n-1) + fib2(s, n-2)
	}
	return s[n]
}

//动态规划(自底向上)
func fib3(n int) int {
	if n <= 0 {
		return n
	}
	s := make([]int, n+1)
	s[0] = 0
	s[1] = 1
	for i := 2; i < n+1; i++ {
		s[i] = s[i-1] + s[i-2]
	}
	return s[n]
}

func recordTime(f func()) {
	t := time.Now()
	f()
	println(time.Since(t).Milliseconds())
}
