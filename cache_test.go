package main

import (
	"reflect"
	"testing"
)

var cache *db

func TestSet(t *testing.T) {
	createCache(t)

	type args struct {
		key   string
		value interface{}
	}

	type result struct {
		value interface{}
		err   error
	}

	tests := []struct {
		name    string
		args    args
		results result
	}{
		{
			name: "Success",
			args: args{
				key:   "key",
				value: "value",
			},
			results: result{
				value: "value",
				err:   nil,
			},
		},
		{
			name: "key cannot be empty",
			args: args{
				key:   "",
				value: "value",
			},
			results: result{
				value: nil,
				err:   ErrKeyEmpty,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cache.Set(tt.args.key, tt.args.value)

			if !reflect.DeepEqual(err, tt.results.err) {
				t.Errorf("Expected error to be %v, got %v", tt.results.err, err)
			}
		})
	}
}

func TestGet(t *testing.T) {
	createCache(t)

	type args struct {
		key string
	}

	type result struct {
		value interface{}
		err   error
	}

	tests := []struct {
		name    string
		args    args
		results result
	}{
		{
			name: "Success",
			args: args{
				key: "key",
			},
			results: result{
				value: "value",
				err:   nil,
			},
		},
		{
			name: "key cannot be empty",
			args: args{
				key: "",
			},
			results: result{
				value: nil,
				err:   ErrKeyEmpty,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache.Set(tt.args.key, "value")

			value, err := cache.Get(tt.args.key)

			if !reflect.DeepEqual(err, tt.results.err) {
				t.Errorf("Expected error to be %v, got %v", tt.results.err, err)
			}
			if !reflect.DeepEqual(value, tt.results.value) {
				t.Errorf("Expected value to be %v, got %v", tt.results.value, value)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	createCache(t)

	type args struct {
		key string
	}

	type result struct {
		err error
	}

	tests := []struct {
		name    string
		args    args
		results result
	}{
		{
			name: "Success",
			args: args{
				key: "key",
			},
			results: result{
				err: nil,
			},
		},
		{
			name: "key cannot be empty",
			args: args{
				key: "",
			},
			results: result{
				err: ErrKeyEmpty,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache.Set(tt.args.key, "value")

			err := cache.Delete(tt.args.key)

			if !reflect.DeepEqual(err, tt.results.err) {
				t.Errorf("Expected error to be %v, got %v", tt.results.err, err)
			}
		})
	}
}

func createCache(t *testing.T) {
	cache = New()

	// callback function to reset cache
	t.Cleanup(func() {
		cache = New()
	})
}
