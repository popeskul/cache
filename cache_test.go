package cache

import (
	"testing"
)

var cache *db

func TestSet(t *testing.T) {
	createCache(t, nil)

	cache.Set("key", "value")
	if cache.Get("key") != "value" {
		t.Error("Expected value to be 'value'")
	}
}

func TestGet(t *testing.T) {
	createCache(t, nil)

	cache.Set("key", "value")
	if cache.Get("key") != "value" {
		t.Error("Expected value to be 'value'")
	}
}

func TestDelete(t *testing.T) {
	createCache(t, nil)

	cache.Set("key", "value")
	cache.Delete("key")
	if cache.Get("key") != nil {
		t.Error("Expected value to be nil")
	}
}

func createCache(t *testing.T, data map[string]interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}

	cache = New(data)

	// callback function to reset cache
	t.Cleanup(func() {
		cache = New(map[string]interface{}{})
	})
}
