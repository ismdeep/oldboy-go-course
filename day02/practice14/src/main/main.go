package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int())
	}
	fmt.Println()

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}
	fmt.Println()

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Float64())
	}
	fmt.Println()

	str1 := "hello"
	str2 := "world"
	//str3 := str1 + " " + str2
	str4 := fmt.Sprintf("%s %s", str1, str2)
	n := len(str4)
	str5 := str4[4:5]
	fmt.Println(str5)
	println(n)
	tmp := []byte(str4)
	fmt.Println(len(tmp))
}
