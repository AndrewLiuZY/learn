package iteration

//Repeat return a string of repeat 's' string n times
func Repeat(s string,n int)  string {
	var result string

	for i := 0; i < n; i++ {
		result += s
	}

	return result
}
