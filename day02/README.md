Day02

----

### 文件名、关键字、标识符

1. 所有 go 源码以 .go 结尾

2. 标识符以字母或下划线开头，大小写敏感。

3. _ 是特殊标识符，用来忽略结果。

4. 保留关键字

    - break

    - default

    - func

    - interface

    - select

    - case

    - defer

    - go

    - map

    - struct

    - chan

    - else

    - goto

    - package

    - switch

    - const

    - fallthough

    - if

    - range

    - type

    - continue

    - for

    - import

        ```go
        import "fmt"
        import "os"
        ```

        ```go
        import (
        	"fmt"
        	"os"
        )
        ```

    - return

    - var

5. Go 程序的基本结构

    ```
    package main
    
    import "fmt"
    
    func main() {
    	fmt.Println("Hello, world.")
    }
    ```

    - 任何一个代码文件都隶属于一个包

    - import 关键字，引用其他包

        ```go
        import "fmt"
        import "os"
        ```

        ```go
        import (
        	"fmt"
        	"os"
        )
        ```

    - golang 可执行程序，`package main` ，并且只有一个 `main` 入口函数。

    - 包的函数调用

        - 同一个包中的函数，直接调用
        - 不同包中的函数通过 包+点+函数名 进行调用。如：`fmt.Println()`

    - 包访问控制规则

        - 大写意味着这个函数、变量是可导出的。
        - 小写意味着这个函数、变量是私有的，包外部不可访问。

6. 包的使用

    - Go 的每个源文件都可以包含一个 init 函数
    - 包可以只作初始化，不引用 `import ( _ "add")`

7. 函数的声明与注释

    - 函数的声明：`func 函数名 (参数列表) (返回值列表) {}`

        举例：

        ```
        func add() {
        }
        
        func add(a int, b int) int {
        }
        
        func add(a int, b int) (int, int) {
        }
        ```

        

    - 注释，两种注释：单行注释：`//` 多行注释：`/**/`

        举例：

        ```
        // 计算两个整数的和并返回结果
        func add(a int, b int) int {
        }
        
        /* 计算两个整数的和并返回结果 */
        func add(a int, b int) int {
        }
        ```

        

8. 常量

    - 常量使用 `const` 进行修饰，代表永远是只读的，不可修改。

    - `const` 只能修饰 `boolean` , number ( `int` 相关类型，浮点类型，`complex` ) 和 `string`

    - 语法 `const identifier [type] = value`, 其中 type可以省略。

        举例：

        ```
        const b string = "hello"
        const b = "hello"
        const Pi = 3.141592653
        ```

    - 比较优雅的写法

        ```
        const (
        	a = 1
        	b = 2
        	name = "ismdeep"
        	Pi = 3.141592653
        )
        ```

    - 更加专业的写法

        ```
        const (
        	a = iota
        	b
        	c
        	d
        	e
        )
        ```

9. 变量

    - 语法： `var identifier type`

        举例：

        ```
        var a int
        var b string
        var c string = "hello"
        ```

        举例2（高端写法）：

        ```
        var (
        	a int
        	b int
        	c = 8
        	d string = "ismdeep"
        )
        ```

10. 值类型和引用类型

    - 值类型：变量直接存储值，内存通常在栈中分配。如：基本数据类型 `int` 、 `float` 、 `bool` 、 `string` 以及数组和 `struct`
    - 引用类型：变量存储的是一个地址，这个地址指向最终的值。内存通常在堆上分配。通过 GC 回收。如：指针 、`slice` 、`map` 、`chan`

11. 变量的作用域

    - 在函数内部声明的变量叫做局部变量，生命周期仅限于函数内部。
    - 在函数外部声明的变量叫做全局变量，生命周期作用于整个包，如果是大写的，则作用于整个程序。

12. 数据类型和操作符

    - bool 类型只能存 `true` 和 `false`

    - 相关操作符 `! && ||`

    - 数字类型：`int int8 int16 int32 int64 uint8 uint16 uint32 uint64 float32 float64`

    - 类型转换：`type(variable)` ，比如：

        ```
        var a int = 8
        var b int32 = int32(a)
        ```

        不同类型之间赋值必须使用类型转换，即使是 `int` 和 `int32` 之间也不行。

    - 逻辑操作符：`== != < <= > >=`

    - 数学操作符：`+ - * /`

    - 字符类型：`var a byte = 'C'`

    - 字符串类型：`var a string = "ismdeep"`

        字符串的两种表示方式：双引号，反引号。双引号接受字符转义，反引号不接受字符转义且可多行编辑。

    -  字符串操作

        ```
        str1 := "hello"
        str2 := "world"
        str3 := str1 + " " + str2 // 字符串拼接
        str4 := fmt.Sprintf("%s %s", str1, str2) // 通过字符串格式化进行拼接
        n := len(str4) // 获取字符串长度
        str5 := str4[4:5] // 字符串切片
        str6 := str4[1:] // 字符串切片
        tmp := []byte(str4) // 字符串转字符数组
        ```

13. 一个面试题：当在浏览器中输入 `www.baidu.com` 会发生什么？

    

