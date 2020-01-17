package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start server ....")
	listen, err := net.Listen("tcp", "0.0.0.0:50000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept failed, err:", err)
			continue
		}
		fmt.Println(conn.RemoteAddr())
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read err:", err)
			return
		}
		fmt.Println("Read:", string(buf))
	}
}