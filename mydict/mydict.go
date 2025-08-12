package mydict

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("could not find the word you were looking for")
var errWordExits = errors.New("word already exists")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, definition string) error {
	// we don't need the * on the receiver because maps on Go are automatically using *
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = definition
		return nil
	}
	return errWordExits
}
