package db

import (
	"errors"
)

type DB[T any] struct {
	data map[int]T
}

func NewDB[T any]() *DB[T] {
	return &DB[T]{
		data: make(map[int]T),
	}
}

func (db *DB[T]) Create(id int, item T) error {
	if _, exists := db.data[id]; exists {
		return errors.New("item already exists")
	}
	db.data[id] = item
	return nil
}

func (db *DB[T]) GetByID(id int) (T, error) {
	item, exists := db.data[id]
	if !exists {
		var zero T
		return zero, errors.New("item not found")
	}
	return item, nil
}

func (db *DB[T]) GetAll() ([]T, error) {
	var result []T
	for _, item := range db.data {
		result = append(result, item)
	}
	return result, nil
}

func (db *DB[T]) Update(id int, item T) error {
	if _, exists := db.data[id]; !exists {
		return errors.New("item not found")
	}
	db.data[id] = item
	return nil
}

func (db *DB[T]) Delete(id int) error {
	if _, exists := db.data[id]; !exists {
		return errors.New("item not found")
	}
	delete(db.data, id)
	return nil
}
