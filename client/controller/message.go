package controller

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func ReadMessages(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println(message)
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading from server:", scanner.Err())
	}
}

func WriteMessages(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading from stdin:", scanner.Err())
	}
}
