package main

const (
	ErrNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrWordExists        = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExists = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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
func (d Dictionary) Update(word, meaning string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = meaning
	default:
		return err
	}
	return nil
}
