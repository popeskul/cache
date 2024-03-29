package cache

import (
	"testing"
	"time"
)

var cache *Cache

func TestSet(t *testing.T) {
	createCache(t)

	t.Run("Get", func(t *testing.T) {
		cache.Set("key", "value")

		result, ok := cache.Get("key")
		if !ok {
			t.Error("Expected value to be 'value'")
		}
		if result != "value" {
			t.Error("Expected value to be 'value'")
		}
	})

	t.Run("wait for the key to expire", func(t *testing.T) {
		t.Parallel()

		cache.Set("key", "value")

		time.Sleep(time.Second * 2)

		result, ok := cache.Get("key")
		if ok {
			t.Error("Expected value to be nil")
		}
		if result != nil {
			t.Error("Expected value to be nil")
		}
	})
}

func TestGet(t *testing.T) {
	createCache(t)

	cache.Set("key", "value")

	result, ok := cache.Get("key")
	if !ok {
		t.Error("Expected value to be 'value'")
	}
	if result != "value" {
		t.Error("Expected value to be 'value'")
	}
}

func TestDelete(t *testing.T) {
	createCache(t)

	cache.Set("key", "value")
	cache.Delete("key")

	result, ok := cache.Get("key")
	if ok {
		t.Error("Expected value to be nil")
	}
	if result != nil {
		t.Error("Expected value to be nil")
	}
}

func createCache(t *testing.T) {
	cache = New(time.Millisecond * 1)

	// callback function to reset cache
	t.Cleanup(func() {
		cache = New(time.Millisecond * 1)
	})
}
