package main

import (
	"reflect"
	"testing"
	"testing/fstest"
)

const (
	firstBody = `Title: Post 1
Description: Description 1
Tags: TDD go
---
Hello
 world`
	secondBody = `Title: Post 2
Description: Description 2
Tags : SleepyCode
---
Hello`
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	got := posts[0]
	want := Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"TDD", "go"},
		Body: `Hello
 world`,
	}
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	assertPost(t, got, want)
}
func assertPost(t *testing.T, got, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
