package practice

import "errors"

type Dictionary map[string]string

var (
	ErrWordAlreadyExists = errors.New("word already exists")
	ErrWordNotFound      = errors.New("could not found word")
)

func (d Dictionary) Search(word string) (string, error) {
	foundWord, ok := d[word]

	if !ok {
		return "", ErrWordNotFound
	}

	return foundWord, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		d[word] = definition

	case nil:
		return ErrWordAlreadyExists

	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	if err != nil {
		return err
	}

	d[word] = definition

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
