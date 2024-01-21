package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type Result struct {
	JSON struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"json"`
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

func main() {
	user := User{
		Name: "John",
		ID:   1,
	}
	postRequest, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
		return
	}

	response, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(postRequest))
	if err != nil {
		log.Fatal(err)
		return
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var jsonResult Result
	err = json.Unmarshal(data, &jsonResult)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%+v\n", jsonResult)
}
