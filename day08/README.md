Day 08

---

### Goroutine

1. 进程和线程

   - 进程是程序在操作系统中的一次执行过程，系统进行资源分配和调度的一个独立单位。
   - 线程是进程的一个执行实体，是 CPU 调度和分配的基本单位，它是比进程更小的能独立运行的基本单位。
   - 一个进程可以创建和撤销多个线程，同一个进程中的多个线程之间可以并发执行。

2. 并发和并行

   - 多线程程序在一个核的CPU上运行，就是并发。
   - 多线程程序在多个核的CPU上运行，就是并行。

3. 协程和线程

   - 协程：独立的栈空间，共享堆空间，调度由用户自己控制，本质上有点类似于用户级线程，这些用户级线程的调度也是自己实现的。
   - 线程：一个线程上可以跑多个协程，协程是轻量级的线程。

4. Goroutine 调度模型

   M(操作系统级别) P(线程) G(协程)

5. 如何设置 Go 运行的 CPU 核数

   ```go
   package main
   
   import (
   	"fmt"
   	"runtime"
   )
   
   func main() {
   	num := runtime.NumCPU()
   	runtime.GOMAXPROCS(num)
   	fmt.Println(num)
   }
   ```



### Channel

1. 不同 Goroutine 之间如何通讯？

   - 全局变量和锁同步
   - Channel

2. channel 概念

   - 类似 unix 中的管道（PIPE）
   - 先进先出
   - 线程安全，多个 Goroutine 同时访问，不需要加锁
   - Channel 是有类型的，一个整数的 Channel 只能存放整数。

3. Channel 声明

   - `var test chan int`
   - `var test chan *Student`

4. Channel 初始化

   使用 make 进行初始化。

   ```go
   var test chan int
   test = make(chan int, 10)
   ```

   ```go
   var test chan string
   test = make(chan string, 10)
   ```

5. Channel 基本操作

   ```go
   package main
   
   import (
   	"fmt"
   )
   
   func main() {
   	var testChan chan int
   	testChan = make(chan int, 10)
   
   	/* 往 Channel 中写入数据 */
   	testChan <- 1
   
   	var a int
   
   	/* 从 Channel 中读取数据 */
   	a = <- testChan
   	
   	fmt.Println(a)
   }
   ```

6. Goroutine 和 Channel 相结合

   ```go
   package main
   
   import (
   	"fmt"
   	"time"
   )
   
   func write(ch chan int) {
   	for i := 0; i < 100; i++ {
   		ch <- i
   	}
   }
   
   func read(ch chan int) {
   	for {
   		fmt.Println(<- ch)
   	}
   }
   
   func main() {
   	intChan := make(chan int, 10)
   	go write(intChan)
   	go read(intChan)
   
   	time.Sleep(10 * time.Second)
   }
   ```

7. Channel 阻塞

8.  带缓冲区的 Channel

9. Channel 之间的同步

10. for range 遍历 Channel

11. Channel 的关闭

12. Channel 的只读和只写

13. 对 Channel 进行 select 操作

14. 定时器的使用

15. 一次定时器

16. 超时控制

17. Goroutine 中使用 recover

18. 。。。

### 单元测试

1. 文件必须以 `_test.go` 结尾
2. 使用 `go test` 进行测试