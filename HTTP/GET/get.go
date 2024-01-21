package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("http://golang.org")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}