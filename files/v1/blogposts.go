package main

import (
	"errors"
	"io"
	"io/fs"
	"testing/fstest"
)

type Post struct {
	Title string
}

func NewPostsFromFS(fileSystem fstest.MapFS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return posts, nil
		}
		posts = append(posts, post)
	}
	return posts, nil
}

type stubFailingFS struct {
}

func (s stubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("Failstub")
}

func getPost(filesSystem fs.FS, fileName string) (Post, error) {
	postFile, err := filesSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return readFile(postFile)
}

func readFile(postFile io.Reader) (Post, error) {

	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}
	post := Post{Title: string(postData)[7:]}
	return post, nil
}
