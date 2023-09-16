package main

import (
	"elearning/client/controller"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Chat Client started...")
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	go controller.ReadMessages(conn)
	controller.WriteMessages(conn)
}
