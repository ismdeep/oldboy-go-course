package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scanf("%s", &str)
	n, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("can not convert to int\n")
		fmt.Println(err)
	} else {
		fmt.Printf("Input value is: %d\n", n)
	}
}
