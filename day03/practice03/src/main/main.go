package main

import (
	"fmt"
	"strings"
)

func StrIndex(str string, substr string) (int, int) {
	return strings.Index(str, substr), strings.LastIndex(str, substr)
}

func main() {
	var (
		str string
		substr string
	)
	fmt.Printf("Please Input Str: ")
	fmt.Scanf("%s", &str)
	fmt.Printf("Please Input Substr: ")
	fmt.Scanf("%s", &substr)
	fmt.Println(StrIndex(str, substr))
}

