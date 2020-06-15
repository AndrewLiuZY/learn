package main

func Sum(arr []int) int {
	sum := 0

	for _, num := range arr {
		sum += num
	}

	return sum
}

func SumAll(arrs ...[]int) []int {

	result := make([]int, len(arrs))

	for i, arr := range arrs {
		result[i] = Sum(arr)
	}

	return result
}

func SumAllTails(arrs ...[]int) []int {

	var result []int

	for _, arr := range arrs {
		if len(arr) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(arr[1:]))
		}
	}

	return result
}
