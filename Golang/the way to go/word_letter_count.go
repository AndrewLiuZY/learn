package main

import "bufio"

import "os"

import "fmt"

import "strings"

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('S')
	if err != nil {
		fmt.Println("There were errors reading.exiting program")
	}
	fmt.Println("length:", len(strings.ReplaceAll(input, "\n", "")))
	fmt.Println("word:", strings.Split(input, " "))
	fmt.Println("rowNum:", strings.Count(input, "\n")+1)
}
