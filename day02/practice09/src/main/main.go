package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	var (
		a = 3
		b = 4
	)
	fmt.Printf("a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
}