package main

import "fmt"

func main() {
	fmt.Println(Hello("hello", ""))
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case chinese:
		prefix = chinesePrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

const chinese = "Chinese"
const french = "French"
const englishPrefix = "Hello "
const chinesePrefix = "你好 "
const frenchPrefix = "Bonjour "
