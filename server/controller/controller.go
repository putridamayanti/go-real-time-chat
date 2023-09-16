package controller

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

var (
	clients     = make(map[net.Conn]struct{})
	clientsLock sync.Mutex
)

func HandleClient(conn net.Conn)  {
	fmt.Printf("New connection from %s\n", conn.RemoteAddr())

	clientsLock.Lock()
	clients[conn] = struct{}{}
	clientsLock.Unlock()

	defer func() {
		conn.Close()
		clientsLock.Lock()
		delete(clients, conn)
		clientsLock.Unlock()
		fmt.Printf("%s disconnected\n", conn.RemoteAddr())
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("%s: %s\n", conn.RemoteAddr(), message)

		// Broadcast the message to all connected clients
		clientsLock.Lock()
		for client := range clients {
			if client != conn {
				_, err := client.Write([]byte(message + "\n"))
				if err != nil {
					fmt.Println("Error writing to client:", err)
				}
			}
		}
		clientsLock.Unlock()
	}

	if scanner.Err() != nil {
		fmt.Printf("Error reading from %s: %s\n", conn.RemoteAddr(), scanner.Err())
	}
}
