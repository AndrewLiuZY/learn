package main

func main() {

}

//Oddven 是否为奇数
func Oddven(num int) bool {
	return (num / 2) != 0
}

//Top100OfArrayIsOddven 前一百的整数是否为偶数
func Top100OfArrayIsOddven(arr []int) []int {
	s := make([]int, 1)
	i := 0
	for _, item := range arr {
		if !Oddven(item) {
			s = append(s, item)
		}
		if i++; i == 100 {
			break
		}
	}
	return s
}
