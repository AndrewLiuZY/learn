package main

import "fmt"

type IntVector []int

func (v IntVector) sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func main() {
	fmt.Println(IntVector{1, 2, 3}.sum()) // 输出是6
}
