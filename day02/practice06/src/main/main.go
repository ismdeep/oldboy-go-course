package main

import (
	"fmt"
	"time"
)

const (
	Male = 1
	Female = 2
)

func main() {
	if time.Now().Unix() % Female == 0 {
		fmt.Println("Female")
	} else {
		fmt.Println("Male")
	}
}
