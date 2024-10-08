package myDict

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("not found")
var errWordExist = errors.New("thats word already exist")
var errCantUpdate = errors.New("Cant update non-existing word")

func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]

	if exist {
		return value, nil
	}

	return "", errNotFound
}

func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExist
	}

	return nil
}

func (d Dictionary) Update(word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
