package main

import "fmt"


func add (a int, b int) (sum int, avg int) {
	sum = a + b
	avg = sum / 2
	return sum, avg
}

func modify(a *int) {
	*a = 100
}

func main() {
	a := 8
	fmt.Println(a)
	modify(&a)
	fmt.Println(a)
}
