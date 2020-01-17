package main

import "fmt"

func isPalindromeStr(str string) bool {
	stop := len(str) / 2
	for i := 0; i <= stop; i++ {
		if str[i] != str[len(str) - i - 1] {
			return false
		}
	}
	return true
}

func main() {
	var str string
	fmt.Scanf("%s", &str)
	if isPalindromeStr(str) {
		fmt.Printf("%s is a palindrome string.\n", str)
	} else {
		fmt.Printf("%s is not a palindrome string.\n", str)
	}
}