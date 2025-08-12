package mydict

import "errors"

type Dictionary map[string]string

var (
	errNotFound   = errors.New("could not find the word you were looking for")
	errCantUpdate = errors.New("could not update non-existent word")
	errWordExits  = errors.New("word already exists")
)

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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err != nil {
		return err
	}
	delete(d, word)
	return nil
}
