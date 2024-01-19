package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		var source string
		fmt.Print("Client: ")
		_, err := fmt.Scan(&source)
		if err != nil {
			fmt.Println("Error")
			continue
		}
		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			fmt.Println("Error writing to the server:", err)
			break
		}
		response := make([]byte, 1024)
		n, err := conn.Read(response)
		if err != nil {
			fmt.Println("Error reading from the server:", err)
			break
		}
		fmt.Println("Server:", string(response[:n]))
	}
}
