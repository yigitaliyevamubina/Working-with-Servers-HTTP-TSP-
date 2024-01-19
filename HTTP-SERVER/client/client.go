package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	for {
		var message string
		fmt.Print("Input something to the server: ")
		fmt.Scan(&message)

		response, err := http.Post("http://127.0.0.1:8080", "text/plain", bytes.NewBufferString(message))
		if err != nil {
			fmt.Println("Error while posting a message to the server: ", err)
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error while receiving a message from the server: ", err)
		}

		fmt.Print("Answer: ")
		fmt.Println(string(body))
		time.Sleep(time.Second)
	}
}
