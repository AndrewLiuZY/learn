package main

import "fmt"

type Number int

func main() {
	var n Number = 123
	fmt.Println(n)
	n.Change(456)
	fmt.Println(n)
}

func (n *Number) Change(num int) {
	*n = Number (num)
}
