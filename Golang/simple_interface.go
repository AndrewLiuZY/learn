package main

import "fmt"

func main() {
	var simple Simple
	var rsimple RSimple
	fi(&simple)
	fi(&rsimple)
}

//Simpler a interface have method Get() and Set()
type Simpler interface {
	Get() int
	Set(i int)
}

func fi(s Simpler)  {
	switch s.(type) {
	case *RSimple:
		fmt.Println("type is *RSimple")
	case *Simple:
		fmt.Println("type is *Simple")
		
	}
}

//Simple asd
type Simple int

//RSimple sadas
type RSimple int

//Get asdasd
func (s RSimple) Get() int {
	return int(s+1)
}

//Set a
func (s *RSimple) Set(i int) {
	*s = RSimple(i)
}

//Get sds
func (s Simple) Get() int {
	return int(s)
}

//Set a
func (s *Simple) Set(i int) {
	*s = Simple(i)
}

//Call a
func Call(s Simpler, i int) int {
	s.Set(i)
	return s.Get()
}
