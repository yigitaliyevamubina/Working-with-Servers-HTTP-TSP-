package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	URL := "https://example.com/api/resource"
	params := url.Values{}
	params.Add("param1", "value1")
	params.Add("param2", "value2")
	fullUrl := URL + "?" + params.Encode()

	response, err := http.Get(fullUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

}
