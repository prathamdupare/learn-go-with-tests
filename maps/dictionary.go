package main

import "errors"

type Dictionary map[string]string

var m map[string]string
var ErrNotFound = errors.New("could not find the word you were looking for")

var dictionary = make(map[string]string)

func (d Dictionary) Add(word string) (string, error) {
	return d[word], nil
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
