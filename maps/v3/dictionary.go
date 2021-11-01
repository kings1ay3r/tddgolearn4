package main

import "errors"

var ErrNotFound = errors.New("word-not-found")
var ErrWordExists = errors.New("word-exists")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	if d[word] == "" {
		return "", ErrNotFound
	}
	return d[word], nil
}
func (d Dictionary) Add(word, meaning string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = meaning
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
