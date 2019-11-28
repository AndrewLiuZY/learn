package main

import (
	"container/list"
	"fmt"
)

func main() {
	list := list.New().Init()
	list.PushFront(101)
	list.PushFront(102)
	list.PushFront(103)
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
