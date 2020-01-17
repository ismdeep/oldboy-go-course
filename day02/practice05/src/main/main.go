/* 练习5， 包只作初始化，不引用。 */
package main

import (
	_ "add"
	"fmt"
)

func main() {
	fmt.Println("add not refer:")
}
