Day 06

---

1. 定义

    `Interface` 类型可以定义一组方法，但是这些不需要实现。并且 `interface` 不能包含任何变量。

    接口是用来定义规范的。

2. 定义示例

    ```go
    type example interface {
        Method1(参数列表) 返回值列表
        Method2(参数列表) 返回值列表
    }
    ```

3. `Interface` 类型默认是一个指针

    ```go
    type example interface {
        Method1(参数列表) 返回值列表
        Method2(参数列表) 返回值列表
        ...
    }
    ```

    ```go
    var a example
    a.Method1()
    ```

4. 接口实现

    - Go 中的接口，不需要显示的实现，只要一个变量含有接口类型中的所有方法，那么这个变量就实现了这个接口。因此，Go 中没有 implement 类似的关键字。
    - 如果一个变量含有了多个 interface 类型的方法，那么这个变量就实现了多个接口。
    - 如果一个变量只含有 1 个 interface 的部分方法，那么这个变量就没有实现这个接口。

5. 多态

    一种事物的多种形态，都可以按照统一的接口进行操作。

6. 接口嵌套

    一个接口可以嵌套在另外的接口里面。

7. 类型断言，由于接口是一般类型，不知道具体类型，如果要转成具体类型，可以采用以下方法进行转换。

    ```go
    var t int
    var x interface{}
    x = t
    y = x.(int)
    ```

8. 。。。