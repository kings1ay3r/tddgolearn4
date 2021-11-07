package main

import (
	"log"
	"os"
)

func main() {
	posts, err := NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
