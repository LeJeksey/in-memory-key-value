package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	tt := []struct {
		name  string
		key   string
		value string
	}{
		{name: "Set key1", key: "key1", value: "value1"},
		{name: "Set key2", key: "key2", value: "value2"},
		{name: "Set key3", key: "key3", value: "value3"},
	}

	s := NewStorage()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_ = s.Set(tc.key, tc.value)
			assert.Equal(t, tc.value, s.hash[tc.key])
		})
	}
}

func TestGet(t *testing.T) {
	tt := []struct {
		name  string
		key   string
		value string
	}{
		{name: "Get key1", key: "key1", value: "value1"},
		{name: "Get key2", key: "key2", value: "value2"},
		{name: "Get key3", key: "key3", value: "value3"},
	}

	s := NewStorage()
	for _, tc := range tt {
		s.hash[tc.key] = tc.value
		t.Run(tc.name, func(t *testing.T) {
			res, _ := s.Get(tc.key)
			assert.Equal(t, tc.value, res)
		})
	}
}

func TestDelete(t *testing.T) {
	tt := []struct {
		name  string
		key   string
		value string
	}{
		{name: "Delete key1", key: "key1", value: "value1"},
		{name: "Delete key2", key: "key2", value: "value2"},
		{name: "Delete key3", key: "key3", value: "value3"},
	}

	s := NewStorage()
	for _, tc := range tt {
		s.hash[tc.key] = tc.value
		t.Run(tc.name, func(t *testing.T) {
			_ = s.Delete(tc.key)
			assert.Equal(t, "", s.hash[tc.key])
		})
	}
}
