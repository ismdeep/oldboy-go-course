package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	pool := &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "localhost:6379")
		},
		TestOnBorrow:    nil,
		MaxIdle:         16,
		MaxActive:       0,
		IdleTimeout:     300,
		Wait:            false,
		MaxConnLifetime: 0,
	}
	conn := pool.Get()
	fmt.Println(conn)
}
