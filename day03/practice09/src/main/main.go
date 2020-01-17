package main

import "fmt"

func modify(a *int) {
	*a = 100
}

func main() {
	var a int = 10
	fmt.Printf("a: %d\n", a)
	modify(&a)
	fmt.Printf("a: %d\n", a)
}