package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	// 1. 监听端口
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [512]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("接收到客户端的信息：%s\n", string(buf[:n]))
	}
}
