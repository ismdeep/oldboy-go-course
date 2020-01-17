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
	target := rand.Intn(101)
	count := 0
	cur := -1
	for cur != target {
		fmt.Scanf("%d", &cur)
		count++
		if cur == target {
			fmt.Println("恭喜你，猜对了。")
		} else if cur < target {
			fmt.Println("猜小了。")
		} else {
			fmt.Println("猜大了。")
		}
	}
	fmt.Printf("你一共猜了：%d次\n", count)
}