package main

import "errors"

var ErrNotFound = errors.New("word-not-found")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	if d[word] == "" {
		return "", ErrNotFound
	}
	return d[word], nil
}
