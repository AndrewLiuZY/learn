package main

import "fmt"

func main() {
	s := new(S)
	s.a = 1
	s.b = 2
	s.int = 123
	s.string = "heihei"
	fmt.Println(s)
}

//S a test struct
type S struct {
	a float32
	b float32
	int
	string
}
