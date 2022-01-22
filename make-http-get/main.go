package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PostResponse struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.cypress.io/todos/11")
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()

	var post PostResponse

	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		print(err)
	}

	fmt.Printf("UserId: %v\n", post.UserId)
	fmt.Printf("Id: %v\n", post.Id)
	fmt.Printf("Title: %v\n", post.Title)
	fmt.Printf("Completed: %v\n", post.Completed)
}
