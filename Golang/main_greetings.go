package main

import "time"

import "fmt"

func main() {
	Greet()
}

//ISAM 判断是否为早上
func ISAM() bool {
	return time.Now().Hour() < 12
}

//ISEvening 判断是否为晚上
func ISEvening() bool {
	return time.Now().Hour() > 18
}

//ISAfternoon 判断是否为下午
func ISAfternoon() bool {
	hour := time.Now().Hour()
	return hour >= 12 && hour <= 18
}

//Greet Just a greeting
func Greet() {
	switch {
	case ISAM():
		fmt.Println("Good Morning")
	case ISAfternoon():
		fmt.Println("Good Afternoon")
	case ISEvening():
		fmt.Println("Good Evening")
	}
}
