package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	fmt.Println(Int32FromInt64(-1 << 35))
	fmt.Println(Int32FromInt64(2323))
	fmt.Println(Int32FromInt64(-1 << 37))
	fmt.Println(Int32FromInt64(123123123))
	fmt.Println(Int32FromInt64(123123123123))
}

//ConvertInt64ToInt ConvertInt64ToInt
func ConvertInt64ToInt32(num int64) int32 {
	if num >= math.MinInt32 && num <= math.MaxInt32 {
		return int32(num)
	}
	panic(fmt.Errorf("can't convert this to int32 %d", num))
}

func Int32FromInt64(num int64) int32 {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	return ConvertInt64ToInt32(num)
}
