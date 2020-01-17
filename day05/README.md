Day 05

---

### Go 中的 struct

1. 用来自定义复杂数据结构。
2. `struct` 里面可以包含多个字段（属性）。
3. `struct` 类型可以定义方法，注意和函数的区别。
4. `struct` 类型是值类型。
5. `struct` 类型可以嵌套。
6. Go 语言没有 `class` 类型，只有 `struct` 类型。

### struct 的定义

1. `struct` 声明

    ```go
    type Student struct {
    	Name string
    	Age int
    }
    ```

2. `struct` 中字段访问，和其他语言一样，使用点。

    ```go
    package main
    
    import "fmt"
    
    type Student struct {
    	Name string
    	Age int
    }
    
    func main() {
    	var stu Student
    	stu.Name = "Del Cooper"
    	stu.Age = 18
    	fmt.Printf("Name:%s    Age:%d\n", stu.Name, stu.Age)
    }
    ```

3. `struct` 定义的三种形式

    ```go
    var stu Student
    ```

    ```go
    var stu *Student = new (Student)
    ```

    ```go
    var stu *Student = &Student{}
    ```

    - 后两种放回的都是指向结构体的指针，访问形式如下：

        ```go
        (*stu).Name
        (*stu).Age
        ```

        也可以使用如下：

        ```go
        stu.Name
        stu.Age
        ```

    - Go 语言中创建的动态内存（即上面后两种创建方式）不需要程序员手动释放内存。Go 语言中通过 GC 自动释放内存。

4. `struct` 的内存布局：`struct` 中的所有字段在内存中是连续的。

5. 链表的定义

    ```go
    type Node struct {
    	Val int
    	Next* Node
    }
    ```

    每个节点包含下一个节点的地址，这样把所有的节点串起来了。通常把表中的第一个节点叫做链表头。

    ```go
    package main
    
    import "fmt"
    
    type Node struct {
    	Val  int
    	Next *Node
    }
    
    type LinkedList struct {
    	Head *Node
    	Tail *Node
    }
    
    /* 链表构造函数 */
    func NewLinkedList () *LinkedList {
    	var list = &LinkedList{
    		Head: nil,
    		Tail: nil,
    	}
    	list.Head = &Node{
    		Val:  0,
    		Next: nil,
    	}
    	list.Tail = list.Head
    	return list
    }
    
    /* 判断链表是否为空 */
    func (list *LinkedList) IsEmpty() bool {
    	if list.Head.Next == nil {
    		return true
    	}
    	return false
    }
    
    /* 在尾部插入 */
    func (list *LinkedList) InsertTail(val int) {
    	var node *Node = &Node{
    		Val:  val,
    		Next: nil,
    	}
    	list.Tail.Next = node
    	list.Tail = list.Tail.Next
    }
    
    /* 在头部插入 */
    func (list *LinkedList) InsertFront(val int) {
    	var node *Node = &Node{
    		Val:  val,
    		Next: nil,
    	}
    	node.Next = list.Head.Next
    	if list.Head.Next == nil {
    		list.Tail = node
    	}
    	list.Head.Next = node
    }
    
    /* 遍历 */
    func (list *LinkedList) Traverse(Op func(node *Node)) {
    	var cur = list.Head
    	for cur.Next != nil {
    		cur = cur.Next
    		Op(cur)
    	}
    }
    
    func main() {
    	var list *LinkedList = NewLinkedList()
    	list.InsertFront(5)
    	list.InsertTail(1)
    	list.InsertTail(2)
    	list.InsertTail(4)
    	list.InsertFront(3)
    	list.Traverse(func(node *Node) {
    		fmt.Printf("%d ", node.Val)
    	})
    	fmt.Printf("\n")
    }
    ```

6. 双链表的定义

    ```go
    type Node struct {
    	Val  int
    	Next *Node
    	Prev *Node
    }
    ```

    如果有两个指针分别指向前一个节点和后一个节点，则称作双链表。

7. 二叉树的定义

    ```go
    type Node struct {
    	Val int
    	left *Node
    	right *Node
    }
    ```

    如果每个节点有两个指针分别来指向左子树和右子树，我们把这样的结构叫做二叉树。

    ```go
    package main
    
    import (
    	"fmt"
    	"math/rand"
    	"time"
    )
    
    /*  */
    type Node struct {
    	Val   int
    	left  *Node
    	right *Node
    }
    
    /* Node 构造函数 */
    func NewNode(val int) *Node {
    	return &Node{
    		Val:   val,
    		left:  nil,
    		right: nil,
    	}
    }
    
    /* 节点随机插入 */
    func (node *Node) InsertByRand() {
    	if node == nil {
    		return
    	}
    	flag := rand.Intn(2)
    	if flag == 0 {
    		if node.left == nil {
    			node.left = NewNode(rand.Intn(100))
    		} else {
    			node.left.InsertByRand()
    		}
    	} else {
    		if node.right == nil {
    			node.right = NewNode(rand.Intn(100))
    		} else {
    			node.right.InsertByRand()
    		}
    	}
    }
    
    /* Node DFS 遍历 */
    func (node *Node) DFSTraverse(Op func(node *Node)) {
    	if node == nil {
    		return
    	}
    	Op(node)
    	node.left.DFSTraverse(Op)
    	node.right.DFSTraverse(Op)
    }
    
    /********************************************************/
    
    type BinTree struct {
    	Root *Node
    }
    
    /* 二叉树构造函数 */
    func NewBinTree() *BinTree {
    	var ans *BinTree = &BinTree{Root: nil}
    	return ans
    }
    
    /* 随机在树中插入一个节点 */
    func (tree *BinTree) InsertByRand() {
    	if tree.Root == nil {
    		tree.Root = &Node{
    			Val:   rand.Intn(100),
    			left:  nil,
    			right: nil,
    		}
    		return
    	}
    	tree.Root.InsertByRand()
    }
    
    /* BinTree DFS 遍历 */
    func (tree *BinTree) DFSTraverse(Op func(node *Node)) {
    	tree.Root.DFSTraverse(Op)
    }
    
    func init() {
    	rand.Seed(time.Now().UnixNano())
    }
    
    func main() {
    	var tree = NewBinTree()
    	for i := 0; i < 10; i++ {
    		tree.InsertByRand()
    	}
    	tree.DFSTraverse(func(node *Node) {
    		fmt.Printf("node: %p    node.left: %p    node.right: %p => %d\n", node, node.left, node.right, node.Val)
    	})
    }
    ```

8. 结构体是用户单独定义的类型，不能和其他类型进行强制转换。

9. Go 里面的 struct 没有构造函数，一般可以使用工厂模式来解决这个问题。

    ```go
    type Node struct {
    	X, Y int
    }
    
    func NewNode(x, y int) *Node {
    	return &Node{
    		X: x,
    		Y: y,
    	}
    }
    ```

10. 可以为 struct 中的每个字段写上一个 tag，这个 tag 可以通过反射的机制获取到，最常用的场景就是 json 序列化和反序列化。

    ```go
    package main
    
    import (
    	"encoding/json"
    	"fmt"
    )
    
    type Node struct {
    	X    int    `json:"x"`
    	Y    int    `json:"y"`
    	Name string `json:"name"`
    }
    
    func NewNode(x, y int, name string) *Node {
    	return &Node{
    		X:    x,
    		Y:    y,
    		Name: name,
    	}
    }
    
    func main() {
    	node := NewNode(1, 2, "ismdeep")
    	data, _ := json.Marshal(node)
    	fmt.Println(string(data))
    }
    ```

    ```text
    {"x":1,"y":2,"name":"ismdeep"}
    ```

11. 结构体中字段可以没有名字，即匿名字段。

12. 匿名字段冲突处理

13. Go 中的方法是作用在特定类型的变量上，因此自定义类型都可以有方法，而不仅仅是struct

    定义：`func (recevier type) methodName (参数列表) (返回值列表){}`

14. 方法的调用

15. 方法和函数的区别

16. 指针 receiver vs 值的 receiver

    本质上和函数的值传递和地址传递是一样的

17. 方法的访问控制，通过大小写控制

18. 继承，如果一个 struct 嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的方法，从而实现了继承。

    ```go
    package main
    
    import "fmt"
    
    type Car struct {
    	Weight int
    	Name string
    }
    
    func (car Car) GetName() string {
    	return car.Name
    }
    
    type Bike struct {
    	Car
    }
    
    func main() {
    	var bike Bike
    	bike.Name = "BIKE"
    	fmt.Println(bike.Weight)
    	fmt.Println(bike.GetName())
    }
    ```

19. 组合和匿名字段，如果一个 struct 嵌套了另一个匿名结构体，那么这个结构可以直接访问匿名结构体的方法，从而实现了继承。如果一个 struct 嵌套了另一个有名结构体，那么这个模式就叫组合。

20. 多重继承，如果一个 struct 嵌套了多个匿名结构体，那么这个结构体可以直接访问多个匿名结构体的方法，从而实现了多重继承。

21. 实现 String()

    如果一个变量实现了 String() 这个方法，那么 `fmt.Println()` 默认会调用这个变量的 `String()` 进行输出。

    ```go
    package main
    
    import "fmt"
    
    type Node struct {
    	Data int
    	Name string
    }
    
    func (node Node) String() string {
    	str := fmt.Sprintf("Name: {%s}    Data: {%d}", node.Name, node.Data)
    	return str
    }
    
    func main() {
    	node := Node{
    		Data: 100,
    		Name: "ismdeep",
    	}
    	fmt.Println(node)
    }
    ```

    

### 接口

1. 定义

    `Interface` 类型可以定义一组方法，但是这些不需要实现。并且 `interface` 不能包含任何变量。

2. 定义

    比如：

    ```go
    type example interface {
      Method1(参数列表) 返回值列表
      Method2(参数列表) 返回值列表
    }
    ```

3. `interface` 类型默认是一个指针

    ```go
    type example interface {
      Method1(参数列表) 返回值列表
      Method2(参数列表) 返回值列表
    }
    var a example
    a.Method1()
    ```

4. 接口实现

    - Golang 中的接口，不需要显示的实现。只要一个变量，含有接口类型中的所有方法，那么这个变量就实现了这个接口。因此，Go 中没有 implement 类似的关键字。
    - 如果一个变量含有多个 interface 类型的方法，那么这个变量就实现了多个接口。
    - 如果一个变量只含有1个 interface 的部分方法，那么这个变量就没有实现这个接口。

5. 多态，一种事物的多种形态，都可以按照统一的接口进行操作。

6. 接口嵌套。一个接口可以嵌套在另外的接口。

7. 类型断言，由于接口是一般类型，不知道具体类型，如果要转成具体类型，可以采用以下方法进行转换。

    ```go
    var t int
    var x interface{}
    x = t
    y = x.(int)
    ```

