package main

import (
	"encoding/json"
	"net/http"
	
)


type Post struct {
	Id     int        `json:"id"`
	Title  string    `json:"Title"`
	Text   string     `json:"Text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{Id: 1, Title: "Title 1", Text: "Text 1"}}
}

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content - type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)

}

func addPost(resp http.ResponseWriter, req *http.Request) {
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return

	}
}