package stack

import (
	"fmt"
	"strconv"
)

//Stack a stack
type Stack struct {
	arr []interface{}
}

//Pop return and remove the top of stack
func (stk *Stack) Pop() interface{} {
	defer fmt.Println(stk)
	if len(stk.arr) == 0 {
		return nil
	}
	str := stk.arr[len(stk.arr)-1]
	stk.arr = stk.arr[:len(stk.arr)-1]
	return str
}

//Push push a element to the top of stack
func (stk *Stack) Push(str interface{}) {
	stk.arr = append(stk.arr, str)
	fmt.Println(stk)
}

//Length return the size of stack
func (stk Stack) Len() int {
	for index, item := range stk.arr {
		if item == "" {
			return index
		}
	}
	return len(stk.arr)
}

func (stk Stack) IsEmpty() bool {
	return len(stk.arr) == 0
}

func (stk Stack) String() string {
	str := ""
	for index, item := range stk.arr {
		var s string
		switch t := item.(type) {
		case int:
			s = strconv.Itoa(int(t))
		case bool:
			s = strconv.FormatBool(bool(t))
		case float32:
			s = strconv.FormatFloat(float64(t), 'f', 6, 64)
		case string:
			s = string(t)
		default:
		}
		str += "[" + strconv.Itoa(index) + ":" + s + "] "

	}
	return str
}
