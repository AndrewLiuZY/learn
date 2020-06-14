package main

import "fmt"

func main() {
	andrew := new(Employee)
	andrew.salary = 10000
	fmt.Println("Andrew's salary:", andrew.salary)
	andrew.giveRaise(50)
	fmt.Println("Andrew's salary:", andrew.salary)
}

//Employee Employee
type Employee struct {
	salary float32
}

func (e *Employee) giveRaise(p float32) {
	fmt.Println("Andrew's salary add", p, "percent")
	e.salary = e.salary * (100 + p) / 100
}
