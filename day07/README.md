Day 07

---

### 终端读写

1. 终端读写

    操作终端相关文件句柄常量

    `os.Stdin` 标准输入

    `os.Stdout` 标准输出

    `os.Stderr` 标准错误输出

2. 终端读写示例

3. 带缓冲区的读写

    ```go
    package main
    
    import (
    	"bufio"
    	"fmt"
    	"os"
    )
    
    var inputReader *bufio.Reader
    var input string
    var err error
    
    func main() {
    	inputReader = bufio.NewReader(os.Stdin)
    	fmt.Print("Please enter some input: ")
    	input, err = inputReader.ReadString('\n')
    	if err == nil {
    		fmt.Printf("The input was: %s\n", input)
    	}
    }
    ```

4. 练习，从终端读取一行字符串，统计英文、数字、空格及其他字符的数量。




### 文件读写

1. `os.File` 封装所有文件相关操作，之前讲的 `os.Stdin` , `os.Stdout` , `os.Stderr` 都是 `*os.File`
    - 打开一个文件进行读操作： `os.Open(name string) (*File, error)`
    - 关闭一个文件： `File.Close()`
2. 文件操作示例
3. 读取整个文件示例
4. 读取压缩文件示例
5. 文件写入
6. 拷贝文件



### 命令行参数

1. `os.Args` 是一个 `string` 的切片，用来存储所有的命令行参数
2. flag 包的使用，用来解析命令行参数。
3. 命令行参数解析

### Json 数据协议

1. Json 数据协议
    - 导入包： `import "encoding/json"`
    - 序列化：`json.Marshal(data.interface{})`
    - 反序列化：`json.UnMarshal(data []byte, v interface{})`
2. 。。。