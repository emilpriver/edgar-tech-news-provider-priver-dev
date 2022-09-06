package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	PriverDevProvider "github.com/emilpriver/edgar-tech-news-provider-priver-dev/provider"
)

var port = "1337"

func handlePosts(w http.ResponseWriter, r *http.Request) {
	posts, err := PriverDevProvider.Posts()
	if err != nil {
		fmt.Print(fmt.Errorf("%w", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(posts)
}

func main() {
	http.HandleFunc("/posts", handlePosts)

	log.Printf("Running on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
