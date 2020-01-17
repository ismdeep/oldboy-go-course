Day 09

---

### TCP 编程

1. 客户端和服务器
2. 服务端的处理流程
   - 监听端口
   - 接收客户端的连接
   - 创建 Goroutine, 处理该连接
3. 客户端的处理流程
   - 建立与服务端的连接
   - 进行数据收发
   - 关闭连接

### Redis

1. Redis

   Redis 是一个开源的高性能 key-value 内存数据库，可以把它当成远程的数据结构。

   支持的 value 类型非常多，比如：string，list，set，hash等等。

   Redis 性能非常高，单机能够达到 15w qps，通常适合做缓存。

2. Redis 的使用

   使用第三方开源的 Redis 库：

   github.com/garyburd/redigo/redis

   ```go
   import (
       "github.com/garyburd/redigo/redis"
   )
   ```

3. 连接 Redis

   ```go
   package main
   ```