Day 10

---

1. http 编程

   - Go 原生支持 http, `import "net/http"`
   - Go 的 http 服务性能和 nginx 比较接近
   - 几行代码就可以实现一个 web 服务

2. 简单的例子

   ```go
   package main
   
   import (
   	"fmt"
   	"net/http"
   )
   
   func Hello(w http.ResponseWriter, r *http.Request) {
   	fmt.Println("handle hello")
   	fmt.Fprintf(w, "hello")
   }
   
   func main() {
   	http.HandleFunc("/", Hello)
   	err := http.ListenAndServe("0.0.0.0:8888", nil)
   	if err != nil {
   		fmt.Println("Http Listen Failed. err:", err)
   	}
   }
   ```

3. http client

   ```go
   package main
   
   import (
   	"fmt"
   	"io/ioutil"
   	"net/http"
   )
   
   func main() {
   	res, err := http.Get("https://ismdeep.com/software")
   	if err != nil {
   		fmt.Println("Get err:", err)
   		return
   	}
   	data, err := ioutil.ReadAll(res.Body)
   	if err != nil {
   		fmt.Println("Get data err:", err)
   		return
   	}
   	fmt.Println(string(data))
   }
   ```

4. http 常见请求方法

   - Get 请求
   - Post 请求
   - Put 请求
   - Delete 请求
   - Head 请求

5. head 请求实例

   ```go
   package main
   
   import (
   	"fmt"
   	"net/http"
   	"sync"
   	"time"
   )
   
   var urls = []string{
   	"https://www.baidu.com",
   	"https://www.google.cn",
   	"https://taobao.com",
   	"http://jw.jxust.edu.cn",
   }
   
   func main() {
   	var wg sync.WaitGroup
   	for _, url := range urls {
   		wg.Add(1)
   		go func(url string) {
   			client := http.Client{
   				Transport:     nil,
   				CheckRedirect: nil,
   				Jar:           nil,
   				Timeout:       2 * time.Second,
   			}
   			resp, err := client.Head(url)
   			if err != nil {
   				fmt.Printf("Head %s failed, err: %v\n", url, err)
   			} else {
   				fmt.Printf("Head successed {%s}, Status: %v, Length: %d\n", url, resp.Status, resp.ContentLength)
   			}
   			wg.Add(-1)
   		}(url)
   	}
   	wg.Wait()
   }
   ```

6. http 常见状态码

   ```go
   http.StatusContinue = 100
   http.StatusOK = 200
   http.StatusFound = 302
   http.StatusBadRequest = 400
   http.StatusUnauthorized = 401
   http.StatusForbidden = 403
   http.StatusNotFound = 404
   http.StatusInternalServerError = 500
   ```

7. 表单处理

8. 模板

   - 替换 {{.字段名}}
   - if-else 流程
   - 循环

### MySQL 编程

1. 连接 MySQL

   ```go
   database, err := sqlx.Open("mysql", "username:[password]@tcp(127.0.0.1:3306)/test")
   ```

2. insert 操作

   ```go
   r, err := Db.Exec("insert into person(username, sex, email) values(?,?,?)", "stu001", "man", "std01@qq.com")
   ```

3. ...

