/* 课堂练习 */
package main

import "fmt"

func addSplit(n int) {
	for i := 0; i <= n; i++ {
		fmt.Printf("%d+%d=%d\n", i, n-i, n)
	}
}


func main() {
	addSplit(5)
}
