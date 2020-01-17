package main

import "fmt"

func concat(str string, arg ...string) string {
	ans := str
	for _, s := range arg {
		ans = ans + s
	}
	return ans
}

func main() {
	fmt.Println(concat("ismdeep", "hello"))
}
