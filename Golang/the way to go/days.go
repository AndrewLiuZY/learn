package main

import "fmt"

func main() {
	var d Day
	d = 3
	fmt.Print(d)
}

var days = [...]string{"MO", "TU", "WE", "TH", "FR", "ST", "SU"}

type Day int

func (d Day) String() string {
	return days[d]
}
