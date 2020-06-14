package main

import "errors"

import "log"

import "fmt"

func main() {
	handler := func(num1, num2 int) {
		num, _ := errorHandler(divideBy)(num1, num2)
		fmt.Println(num)
	}
	handler(12, 3)
	handler(5, 3)
	handler(15, 3)
	handler(12, 0)
	handler(12, 3)
}

func divideBy(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("divideBy - Can't divide by 0")
	}
	return num1 / num2, nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func errorHandler(f func(num1, num2 int) (int, error)) func(int, int) (int, error) {
	return func(num1, num2 int) (int, error) {
		defer func() {
			if err, ok := recover().(error); ok {
				log.Printf("run time panic: %v", err)
			}
		}()
		i, err := f(num1, num2)
		check(err)
		return i, err
	}
}
