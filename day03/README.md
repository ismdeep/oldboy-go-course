Day 03

---

### strings 和 strconv 的使用

- `strings.HasPrefix(str string, prefix string) bool` ：判断字符串 `str` 是否以 `prefix` 开头
- `strings.HasSuffix(str string, suffix string) bool` ：判断字符串 `str` 是否以 `suffix` 结尾
- `strings.Index(str string, substr string) int` ： 判断 `substr` 在 `str` 中首次出现的位置，如果没有出现则返回 `-1`
- `strings.LastIndex(str string, substr string) int` ： 判断 `substr` 在 `str` 中最后出现的位置，如果没有出现则返回 `-1`
- `strings.Replace(str string, old string, new string, n int) string` ：将 `str` 中的 `old` 字符串替换成 `new` 字符串，`n` 表示替换的个数，如果 $n \lt 0$ 则表示替换个数无限制。
- `strings.Count(str string, substr string) int` ：字符串计数，但是 `str` 中出现的 `substr` 并不叠加。例如：`eee` 中出现 `ee` 的次数为 `1` 
- `strings.Repeat(str string, count int) string` ：重复 `count` 次 `str`
- `strings.ToUpper(str string) string` ：字符串转为大写
- `strings.ToLower(str string) string` ：字符串转为小写
- `strings.TrimSpace(str string) string` ：去掉 `str` 字符串中的首尾空白字符
- `strings.Trim(str string, cutset string)` ：去掉 `str` 字符串中首尾在 `cutset` 中出现的字符
- `strings.TrimLeft(str string, cutset string)` ：去掉 `str` 字符串首部在 `cutset` 中出现的字符
- `strings.TrimRight(str string, cutset string)` ：去掉 `str` 字符串尾部在 `cutset` 中出现的字符
- `strings.Fields(str string) []string` ：返回 `str` 字符串空格分割的所有子串的 `slice`
- `strings.Split(str string, sep string) []string` ：返回 `str` 字符串通过 `sep` 分割的所有子串的 `slice`
- `strings.Join(a []string, sep string) string` ：用 `sep` 把 `a` 中的所有元素连接起来
- `strconv.Itoa(n int) string` ：将整数 `n` 转化成字符串
- `strconv.Atoi(s string) (int, error)` ：将字符串 `s` 转化成整数

### 时间和日期类型

- `(t Time) Format(layout string) string` ：时间格式化函数，通过 `layout` 字符串对时间进行格式化。`layout` 字符串使用与其他语言中的不太一样。使用一个时间实例进行表示：`"2006-01-02 15:04:05.000"` 如：`time.Now().Format("2006-01-02 15:04:05.000")`
- `time.Now() Time` ：获取当前时间。
- `(t Time) UnixNano() int64` ：获取时间对应的 Unix 时间戳，单位：纳秒。
- `(t Time) Unix() int64` ：获取时间对应的 Unix 时间戳，单位：秒。

### 指针类型

1. 普通类型，变量存的是值，也叫值类型。
2. 获取变量的地址用 `&` ，比如：`var a int` ，获取 `a` 的地址：`&a`
3. 指针类型，变量存的是一个地址，这个地址指向的内存存的才是值。
4. 获得指针所指向的值使用 `*` ，比如：`var *p int` ，使用 `*p` 获取 `p` 指向的内存的值。

### 流程控制

1. `if-else` 分支

2. `switch` 分支，Go语言中的 `switch` 语句不需要 `break` 进行截断。然而可以通过 `fallthough` 关键字使得进入 `case` 之后执行完继续执行后续 `case` 的语句。

3. `for` 语句

    - for` 语句外面是没有括号的。

        ```go
        for i := 0; i < n; i++ {
        }
        ```

    - Go 语言中没有 `while` 语句，使用 `for` 进行代替。

        ```go
        for i < n {
        }
        ```

    - `for range` 语句

        ```go
        for i, v := range("Hello") {
            fmt.Printf("i: %d    v: %c\n", i, v)
        }
        ```

4. `break` 和 `continue` 语句
5. `goto` 和 `label` 语句

### 函数

1. 函数的声明语法：`func 函数名 (参数列表) [(返回值参数列表)] {}`

    ```go
    func add() {
    }
    ```

    ```go
    func add(a int, b int) {
    }
    ```

    ```go
    func add(a int, b int) int {
    }
    ```

    ```go
    func add(a int, b int) (int, int) {
    }
    ```

    ```go
    func add(a, b int) (int, int) {
    }
    ```

    **注意**：花括号的左括号必须跟在函数声明行末

2. Go 函数的特点：

    - 不支持重载，一个包不允许有两个名字一样的函数。
    - 函数是一等公民，函数也是一种类型，一个函数可以赋值给变量。
    - 匿名函数
    - 多返回值

3. 函数参数的传递方式

    - 值传递
    - 引用传递

    **注意1：** 无论是值传递还是引用传递，传递给函数的都是变量副本，不过，值传递是值的拷贝，引用传递是地址的拷贝。一般来说，地址拷贝更为高效，而值拷贝取决于拷贝对象的大小。对象越大，则值拷贝的效率越低。

    **注意2：** `map` 、`slice` 、`chan` 、指针 、`interface` 默认以引用的方式传递。

4. 命名返回值名称：

    ```go
    func add (a int, b int) (sum int, avg int) {
    	sum = a + b
    	avg = sum / 2
    	return sum, avg
    }
    ```

5. `_` 标识符，用来忽略返回值。

6. 可变参数： 

    ```go
    func sum(arg...int) int {
    	ans := 0
    	for _, val := range arg {
    		ans += val
    	}
    	return ans
    }
    ```

    ```
    func max(a int, arg...int) int {
    	maxVal := a
    	for _, item := range arg {
    		if item > maxVal {
    			maxVal = item
    		}
    	}
    	return maxVal
    }
    ```

7. `defer` 的用途

    - 当函数返回时，执行 `defer` 语句。因此可以用来做资源清理。
    - 多个 `defer` 语句，按先进后出的方式执行。
    - `defer` 语句中的变量，在 `defer` 声明时就决定了。

    用途：

    - 关闭文件句柄
    - 锁资源释放
    - 数据库连接释放

