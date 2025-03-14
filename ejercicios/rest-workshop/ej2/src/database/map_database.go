package database

import "errors"

type db[T any] struct {
	elems map[string]T
}

func NewDatabase[T any]() db[T] {
	return db[T]{
		elems: make(map[string]T),
	}
}

func (d db[T]) Create(id string, elem T) (T, error) {
	if _, found := d.elems[id]; found {
		return elem, errors.New("element already exists")
	}

	d.elems[id] = elem
	return elem, nil
}
