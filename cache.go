package cache

import (
	"sync"
)

type db struct {
	data sync.Map
}

func New() *db {
	return &db{
		data: sync.Map{},
	}
}

func (db *db) Set(key string, value interface{}) {
	db.data.Store(key, value)
}

func (db *db) Get(key string) (result interface{}, ok bool) {
	return db.data.Load(key)
}

func (db *db) Delete(key string) {
	db.data.Delete(key)
}
