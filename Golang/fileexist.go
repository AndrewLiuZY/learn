package main

import "os"

import "fmt"

func main() {
	fmt.Println(FileExist("E:/Github/learn/Golang/exec.go"))
	fmt.Println(FileExist("./exe.go"))
	fmt.Println(FileExist("../learn"))
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
