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
	UserId    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	user := User{
		UserId:    2,
		ID:        2,
		Title:     "Example",
		Completed: true,
	}

	jsonReq, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
		return
	}

	url := "https://jsonplaceholder.typicode.com/posts"

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatal(err)
		return
	}

	request.Header.Set("Content-type", "application/json; charset=UTF-8")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("String:")
	fmt.Println(string(data))

	var respUser User

	err = json.Unmarshal(data, &respUser)
	if err != nil {
		panic(err)
	}

	fmt.Println("Struct:")
	fmt.Printf("%+v\n", respUser)

	fmt.Printf("Status code: %s\n", response.Status)
}
