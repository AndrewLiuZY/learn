package main

import "strconv"

import "fmt"

func main() {
	var c Celsius
	c = 12.11
	fmt.Println(c)
}

type Celsius float64

func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 2, 64) + "°C"
}
