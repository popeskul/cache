package main

import (
	"errors"
	"sync"
)

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrKeyEmpty    = errors.New("key cannot be empty")
)

type db struct {
	data sync.Map
}

func New() *db {
	return &db{
		data: sync.Map{},
	}
}

func (db *db) Set(key string, value interface{}) error {
	if key == "" {
		return ErrKeyEmpty
	}

	db.data.Store(key, value)

	return nil
}

func (db *db) Get(key string) (result interface{}, err error) {
	if key == "" {
		return nil, ErrKeyEmpty
	}

	value, ok := db.data.Load(key)
	if !ok {
		return nil, ErrKeyNotFound
	}

	return value, nil
}

func (db *db) Delete(key string) error {
	if key == "" {
		return ErrKeyEmpty
	}

	_, ok := db.data.Load(key)
	if !ok {
		return ErrKeyNotFound
	}

	db.data.Delete(key)

	return nil
}
