package main

import (
	"sync"
)

type db struct {
	data map[string]interface{}
	sync.Mutex
}

func New(data map[string]interface{}) *db {
	return &db{
		data: data,
	}
}

func (db *db) Set(key string, value interface{}) {
	db.Lock()
	defer db.Unlock()

	db.data[key] = value
}

func (db *db) Get(key string) (result interface{}) {
	db.Lock()
	defer db.Unlock()

	return db.data[key]
}

func (db *db) Delete(key string) {
	db.Lock()
	defer db.Unlock()

	delete(db.data, key)
}
