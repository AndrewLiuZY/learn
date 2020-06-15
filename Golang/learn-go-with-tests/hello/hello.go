package main

import "fmt"

func main() {
	fmt.Println(Hello("hello"))
}

func Hello(name string) string {
	return "Hello " + name
}
