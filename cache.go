package cache

import (
	"sync"
	"time"
)

type value struct {
	value interface{}
	ttl   *time.Time
}

type Cache struct {
	ticker *time.Ticker
	data   sync.Map
}

func New() *Cache {
	db := &Cache{
		ticker: time.NewTicker(time.Second * 1),
		data:   sync.Map{},
	}

	go db.backgroundCacheCleaner()

	return db
}

// background goroutine to clean up expired keys in the cache
func (db *Cache) backgroundCacheCleaner() {
	for {
		<-db.ticker.C
		db.data.Range(func(key, v interface{}) bool {
			vv, ok := v.(*value)
			if !ok {
				return true
			}

			if vv.ttl == nil {
				return true
			}

			if time.Now().After(*vv.ttl) {
				db.data.Delete(key)
			}

			return true
		})
	}
}

func (db *Cache) Set(key string, v interface{}, ttl time.Duration) {
	t := time.Now().Add(ttl)
	db.data.Store(key, &value{v, &t})
}

func (db *Cache) Get(key string) (result interface{}, ok bool) {
	load, ok := db.data.Load(key)
	if !ok {
		return nil, false
	}

	vv, ok := load.(*value)
	if !ok {
		return nil, false
	}

	return vv.value, true
}

func (db *Cache) Delete(key string) {
	db.data.Delete(key)
}
