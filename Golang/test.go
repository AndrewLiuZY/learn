package main

import (
	"fmt"

	"./stack"
)

func main() {
	stk := new(stack.Stack)
	fmt.Println(stk.IsEmpty())
	stk.Push(123)
	fmt.Println(stk.IsEmpty())
	stk.Push("asas")
	stk.Push(12.1)
	stk.Push(false)
	stk.Pop()
	fmt.Println(stk.Pop())
	fmt.Println(stk.IsEmpty())
	fmt.Println(stk.Len())
	fmt.Println()

}
