Day 01

---

### Go 调试器的功能

Go 调试器能够帮助程序员更快的找到程序中出现的问题。在没有调试器的情况下，简单的程序有时候能够很容易的找到问题的原因，但如果业务逻辑很复杂的时候就需要借助调试器单步执行去寻找问题的根源了。



### Go 语言特性

1. 垃圾回收
	- 内存自动回收，再也不需要开发人员管理内存。
	- 开发人员能够专注业务实现，降低了程序员的工作负担。
	- 只需要 new 分配内存，不需要释放操作。 

2. 天然并发
   	- 从语言层面支持并发，非常简单。
   - goroutine，轻量级线程，创建成千上万 goroutine 成为可能。
   - 基于 CSP（Communicating Sequential Process）模型实现。

3. channel
   - 管道，类似 unix/linux 中的pipe
   - 多个goroutine之间通过channel进行通信
   - 支持任何类型
4. 多返回值
   - 一个函数可返回多个值



### 管道的使用

管道的创建

```
var pipe chan int
pipe = make(chan int, 3)
```

```
pipe := make(chan int, 3)
```

存放数据

```
pipe <- 2
```

取出数据

```
val = <- pipe
```

获取管道当前大小

```
len(pipe)
```



### Go 语言类型推导

因为 Go 语言是一种强类型的语言。所以在声明变量的时候需要指名变量的类型。比如 `var n int` ，而 Go 语言中还有类型推导，也就是说可以跳过上面的语句来创建一个变量，如 `n := 3` ，这样就可以知道 `n` 是一个整型变量。



### 在函数中使用管道的两种方式

- 全局变量方式
- 传入参数方式



### Go 代码格式化工具 gofmt

使用命令：

```
gofmt filename.go
```

使用命令（替换原文件）：

```
gofmt -w filename.go
```



### Go 语言中神奇的 _

`_` 这个符号用于忽略掉一些变量，比如一个多返回值函数，但是只需要其中一个或几个结果，比如在使用一个计算两个数的和以及平均值的函数时，有时我们只需要两个数的和，这样在使用的时候可以通过 `_` 屏蔽掉返回结果中不需要的数值。如：`sum, _ = calc(2, 4)` , `_, avg = calc(2,4)`



### 包的概念

1. 和 Python 一样，把相同功能的代码放到一个目录，称之为包
2. 包可以被其他包引用
3. main 包是用来生成可执行文件，每个程序只有一个main包
4. 包的主要用途是提高代码的可复用性



### Go 编译执行

首先需要设置 GOPATH 环境变量

文件目录为：`D:\projects\code-practice\oldboy-go-course\src\day01\hello`

设置 GOPATH 为：`D:\projects\code-practice\oldboy-go-course` 

Windows 上通过指令： `set GOPATH=D:\projects\code-practice\oldboy-go-course`

编译指令：`go build day01/hello`

