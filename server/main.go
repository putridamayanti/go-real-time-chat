package main

import (
	"elearning/server/controller"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Chat Server started...")
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go controller.HandleClient(conn)
	}
}
