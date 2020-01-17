package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		str string
		substr string
	)
	fmt.Scanf("%s %s", &str, &substr)
	fmt.Println(strings.Count(str, substr))
	str = strings.Replace(str, substr, "-", -1)
	fmt.Println(str)
	println(strings.Repeat(str, 2))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
}