/* 每个源文件都可以包含一个init函数，这个 init 函数自动被 go 运行框架调用。 */
package main

import (
	"add"
	"fmt"
)

func main() {
	fmt.Println("Name:", add.Name)
	fmt.Println("Age:", add.Age)
}
