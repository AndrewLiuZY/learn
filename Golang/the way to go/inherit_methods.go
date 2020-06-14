package main

import "fmt"

func main() {
	e := new(Employee1)
	e.SetID(123)
	fmt.Println(e.id)
}

//Base base
type Base struct {
	id int
}

//SetID set the id of base
func (b *Base) SetID(i int) {
	b.id = i
}

//ID get id
func (b Base) ID() int {
	return b.id
}

//Person type
type Person struct {
	Base
	firstName string
	lastname  string
}

//Employee1 type
type Employee1 struct {
	Person
	salary int
}
