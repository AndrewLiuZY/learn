package main

import "unsafe"

import "fmt"

func main() {
	size := unsafe.Sizeof(1)
	fmt.Println(size)
}
