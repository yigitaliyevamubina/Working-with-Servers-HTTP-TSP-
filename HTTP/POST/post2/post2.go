package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func main() {
	formData := url.Values{
		"name":    {"Golang"},
		"surname": {"Org"},
	}

	response, err := http.PostForm("https://httpbin.org/post", formData)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer response.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(result["form"])
	// log.Println(result)
}
