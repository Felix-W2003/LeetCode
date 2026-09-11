package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter message:")
		message, _ := reader.ReadString('\n')
		conn.Write([]byte(message))

		respones, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("Server response:%s", respones)
	}
}
