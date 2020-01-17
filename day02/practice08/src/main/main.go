package main

import "fmt"

func modifyInt (a int) {
	a = 10
}

func modifyIntPointer(a *int)  {
	*a = 10
}

func modifyChanInt (a chan int) {
	a <- 1
	a <- 2
}

func main() {
	var a = 100
	var b chan int = make(chan int, 1)
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	modifyInt(a)
	fmt.Println("a =", a)
	modifyIntPointer(&a)
	fmt.Println("a =", a)
}
