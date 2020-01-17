Day 04

---

### 内置函数

1. `close` ：主要用来关闭 `channel`

2. `len` ：用来求长度，比如：`strings` 、`array` 、`slice` 、`map` 、`channel`

3. `new` ：用来分配内存，主要用来分配值类型内存。比如：`int` 、`struct` 。返回的是指针。

    ```go
    var b *int
    b = new(int)
    *b = 100
    ```

4. `make` ：用来分配内存，主要是用来分配引用类型内存。比如：`map` 、`channel` 、`slice` 等类型。

    ```go
    var c chan int
    c = make(chan int, 4)
    ```

    `make` 和 `new` 的区别：

    ```go
    package main
    
    import "fmt"
    
    func main() {
    	s1 := new([]int)
    	fmt.Println(s1)
    
    	s2 := make([]int, 10)
    	fmt.Println(s2)
    
    	*s1 = make([]int, 5)
    	fmt.Println(*s1)
    }
    ```

5. `append` ：用来追加元素到数组、`slice` 中。

    ```go
    var a []int
    a = append(a, 1)
    ```

6. `panic` 、`recover` ：用来做错误处理。

    ```go
    package main
    
    import (
    	"fmt"
    	"math/rand"
    	"time"
    )
    
    func testRecover(a int, b int) {
    	defer func() {
    		if err := recover(); err != nil {
    			fmt.Println(err)
    		}
    	}()
        
    	c := a / b
    	fmt.Printf("%d / %d = %d\n", a, b, c)
    }
    
    func init() {
    	rand.Seed(time.Now().UnixNano())
    }
    
    func main() {
    	for {
    		testRecover(rand.Intn(1000), rand.Intn(10))
    		time.Sleep(time.Millisecond * 100)
    	}
    }
    ```



###  递归函数

自己调用自己的函数就叫做**递归**。

1.  阶乘

    ```go
    package main
    
    import "fmt"
    
    func calc(n int) int {
    	if n <= 1 {
    		return 1
    	}
    	return calc(n - 1) * n
    }
    
    func main() {
    	fmt.Println(calc(10))
    }
    ```

2. Fibonacci 数列

    ```go
    package main
    
    import "fmt"
    
    func fibonacci(n int) int {
    	if n <= 2 {
    		return 1
    	}
    	return fibonacci(n - 1) + fibonacci(n - 2)
    }
    
    func main() {
    	fmt.Println(fibonacci(3))
    }
    ```

3. 递归的设计原则

    - 一个大的问题能够分解的同等问题的小问题。

    - 定义好终止递归的条件。

### 闭包

**闭包：** 一个函数和与其相关的引用环境组合而成的实体。

闭包的一个例子：

```go
package main

import (
	"fmt"
	"strings"
)

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	func1 := makeSuffixFunc(".bmp")
	func2 := makeSuffixFunc(".jpg")
	fmt.Println(func1("hello"))
	fmt.Println(func2("name"))
}
```



### 数组与切片

1. 数组：是同一种数据类型的固定长度序列。

2. 数组定义：`var a [len]int` ，比如：`var a [5]int` ，数组长度一旦定义是不可变的。

3. 长度是数组类型的一部分，因此：`var a[5] int` 和 `var a [10]int` 是不同的类型。

4. 数组可以通过下标进行访问，下标是从0开始，最后一个元素的下标是：`len - 1` 

    ```go
    for i := 0; i < len(a); i++ {
    	fmt.Println(a[i])
    }
    ```

    ```go
    for _, item := range a {
    	fmt.Println(item)
    }
    ```

    

5. 访问越界，如果下标在数组合法范围之外，则触发访问越界。会 panic

6. 数组是值类型，因此改变副本的值并不会改变本身的值。

7. 数组初始化

    ```go
    var a [5]int = [5]int{1,2,3,4,5}
    ```

    ```go
    // 当初始化使用的数值不够的情况下，其余为 0
    var a [5]int = [5]int{1,2,3,4}
    ```

    ```go
    var a = [5]int{1,2,3,4,5}
    ```

    ```go
    var a = [...]int{1,2,3,4,5}
    ```

    ```go
    var a = [5]string{3:"hello world", 4: "tom"}
    ```

8. 多维数组

    ```go
    var a [4][2]int
    ```

    ```go
    var a [4][2]int = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1,0}}
    ```

    ```go
    var a = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1,0}}
    ```

    ```go
    var a [4][2]int = [...][2]int{{0, 1}, {0, -1}, {1, 0}, {-1,0}}
    ```

    ```go
    var a = [...][...]int{{0, 1}, {0, -1}, {1, 0}, {-1,0}}
    ```

9. 多维数组遍历

    ```go
    package main
    
    import "fmt"
    
    func main() {
    	var a [4][2]int = [...][2]int{{0, 1}, {0, -1}, {1, 0}, {-1,0}}
    	for row, arr := range a {
    		for col, val := range arr {
    			fmt.Println(row, col, val)
    		}
    	}
    }
    ```



### 切片

1. 切片：切片是数组的一个引用，因此切片是引用类型。

2. 切片的长度可以改变，因此，切片是一个可变数组。

3. 切片的遍历方式和数组一样。同样可以使用 `len()` 计算长度。

4. `cap` 可以计算 `slice` 的最大容量。`0 <= len(slice) <= cap(slice)` ，其中 `array` 是 `slice` 引用的数组。

5. 切片的定义：`var 变量名 []类型` ，比如：`var str []string`

   ```go
   var a = [5]int{1,2,3,4,5}
   var sliceA = a[2:4]
   fmt.Println(a)
   fmt.Println(sliceA)
   ```

   ```go
   // 直接包含整个数组
   var sliceA = a[:]
   ```

   ```go
   // 只说明结束为止，开始位置默认为0
   var sliceA = a[:3]
   ```

   ```go
   // 只说明开始位置，结束位置默认为最后
   var sliceA = a[2:]
   ```

6. 切片的内存布局

7. 通过 `make` 来创建切片

   ```go
   var slice []int = make([]int, 10)
   ```

   ```go
   var slice []int = make([]int, 10, 100)
   ```

8. 用 `append` 内置函数操作切片

   ```go
   slice = append(slice, 10)
   ```

   ```go
   // 合并两个切片
   var a []int = []int{1,2,3}
   var b []int = []int{4,5,6}
   var c = append(a, b...)
   ```

9. `for range` 遍历切片

10. 切片 `resize`

11. 切片的拷贝

    `copy` 函数：`copy(dst, src []Type)`

12. `string` 与 `slice`

    `string` 底层就是一个 `byte` 数组。因此，可以可以进行切片操作。

13. `string` 的底层布局

14. `string` 根据 `unicode` 进行拆分：`rune` 函数。`s1 := []rune(s)`

15. 排序和查找操作

    排序操作主要都在 `sort` 包中，导入就可以使用了。

    ```go
    import "sort"
    ```

    `sort.Ints()` 对整数进行排序。

    `sort.Strings()` 对字符串进行排序。

    `sort.Float64s()` 对浮点数进行排序。

    `sort.SearchInts(a []int, x int)` 从数组 `a` 中查找 `x` ，其中 `a` 必须是有序的。

    `sort.SearchStrings(a []string, x string)` 从数组 `a` 中查找 `x` ，其中 `a` 必须是有序的。

    `sort.SearchFloat64s(a []float64, x float64)` 从数组 `a` 中查找 `x` ，其中 `a` 必须是有序的。



### map 数据结构

1. map 简介

   key-value 的数据结构，又叫字典或关联数组。

   - 声明

     ```go
     var map1 map[string]int
     var map1 map[string]string
     var map1 map[int]string
     var map1 map[string]map[string]int
     ```

     声明是不会分配内存的，需要使用 `make` 进行初始化。

     ```go
     var map1 map[string]int
     map1 = make(map[string]int)
     map1 = make(map[string]int, 10) // 第二个参数数字是 map 的容量
     ```

     也可以在声明的时候直接初始化。

     ```go
     var map1 map[string]int = map[string]int{
     	"hello": 1,
     }
     ```

     `map` 中后面的值类型也可以是一个 `map`

     ```go
     package main
     
     import "fmt"
     
     func main() {
     	map1 := make(map[string]map[string]string)
     	map1["key1"] = make(map[string]string)
     	map1["key1"]["key2"] = "hello"
     	map1["key1"]["key3"] = "world"
     	fmt.Println(map1)
     }
     ```

2. `map` 相关操作

3. `map` 是引用类型

4. slice of map

5. `map` 排序

   - 先获取所有的 key，把 key 进行排序
   - 按照排好序的 key 进行遍历
   - map 中的 key 是无序的

6. map 反转

   - 初始化另外一个 map，把 key、value 互换即可。



### 包

1. golang中的包

   - golang目前有150个标准的包，覆盖了几乎所有的基础库
   - golang.org有所有包的文档

2. 线程同步

   - `import ("sync")`
   - 互斥锁 `var mu sync.Mutex`
   - 读写锁 `var mu sync.RWMutex`

3. go get 安装第三方包

   ```bash
   go get github.com/go-sql-driver/mysql
   ```

   