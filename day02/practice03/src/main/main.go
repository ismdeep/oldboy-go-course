/* 练习3，使用包别名 */
package main

import (
	a "add"
	"fmt"
)

func main() {
	fmt.Println("result:", a.Name)
	fmt.Println("result:", a.Age)
}
