package main

import (
	"bufio"
	"errors"
	"io"
	"io/fs"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSperator         = "Tags: "
	bodySeperator        = "---"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
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

	postData, err := newPost(postFile)
	if err != nil {
		return Post{}, err
	}
	// post := Post{Title: string(postData)[7:]}
	return postData, nil
}
func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}
	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}
	readTags := func(tagName string) []string {
		scanner.Scan()
		tagsstring := strings.TrimPrefix(scanner.Text(), tagName)
		return strings.Split(tagsstring, " ")
	}
	readBody := func(bodySeperator string) string {
		for scanner.Text() != bodySeperator {
			scanner.Scan()
		}

		body := ""
		for scanner.Scan() {

			// scanner.Scan()
			body += scanner.Text() + "\n"
		}
		return strings.TrimSuffix(body, "\n")
	}

	titleLine := readLine()[len(titleSeparator):]

	descriptionLine := readMetaLine(descriptionSeparator)

	tags := readTags(tagsSperator)

	body := readBody(bodySeperator)

	return Post{
		Title:       titleLine,
		Description: descriptionLine,
		Tags:        tags,
		Body:        body,
	}, nil
}
