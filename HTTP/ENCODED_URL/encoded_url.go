package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	base, err := url.Parse("https://www.example.com")
	if err != nil {
		log.Fatal(err)
		return
	}

	params := url.Values{}
	params.Add("id", "1")
	params.Add("name", "John")

	base.RawQuery = params.Encode()

	fmt.Printf("Encoded URL: %s", base.String())
}
