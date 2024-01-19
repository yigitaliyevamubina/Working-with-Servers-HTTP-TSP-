package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	
	message, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error while reading a message from the client:", err)
	}
	
	fmt.Print("Received: ")
	fmt.Println(string(message))

	response := []byte("Server is here...")
	w.Write(response)
}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server is listening on port:8080")
	time.Sleep(time.Second)
	fmt.Println("Waiting for a message...")
	http.ListenAndServe(":8080", nil)
}
