package test

import (
	"testing"

	"github.com/Madduxv/mini-redis/internal/storage"
)

func TestHSet(t *testing.T) {
	store := storage.NewStorage()
	defer storage.ClearStorage(store)

	key := "user:1234"
	field := "name"
	value := "Maddux"

	store.HSet(key, field, value)

	if store.StringStore[key][field] != value {
		t.Errorf("HSet failed, expected %v, got %v", value, store.StringStore[key][field])
	}
}

func TestRPush(t *testing.T) {
}

func TestHSetList(t *testing.T) {
	store := storage.NewStorage()
	defer storage.ClearStorage(store)

	key := "user:1234"
	field := "genres"
	value := []string{"ITALIAN", "AMERICAN", "JAPANESE"}

	store.HSetList(key, field, value)

	if len(store.ListStore[key][field]) != len(value) {
		t.Errorf("HSetList failed, expected list of length %d, got %d", len(value), len(store.ListStore[key][field]))
	}

	for i, v := range value {
		if store.ListStore[key][field][i] != v {
			t.Errorf("HSetList failed, expected %v at index %d, got %v", v, i, store.ListStore[key][field][i])
		}
	}
}

func TestHRemove(t *testing.T) {
	store := storage.NewStorage()
	defer storage.ClearStorage(store)

	key := "user:1234"
	field := "name"
	value := "Maddux"

	key1 := "user:1234"
	field1 := "genres"
	value1 := []string{"ITALIAN", "AMERICAN", "JAPANESE"}

	store.HSet(key, field, value)
	store.HSetList(key1, field1, value1)

	store.HRemove(key1)

	_, ok := store.StringStore[key]
	if ok {
		t.Errorf("HRemove failed: Key '%v' still exists in StringStore", key)
	}

	_, ok1 := store.ListStore[key]
	if ok1 {
		t.Errorf("HRemove failed: Key '%v' still exists in ListStore", key)
	}
}

func TestHRemoveStringField(t *testing.T) {
	store := storage.NewStorage()
	defer storage.ClearStorage(store)

	key := "user:1234"
	field := "name"
	value := "Maddux"

	store.HSet(key, field, value)

	store.HRemoveStringField(key, field)

	val, ok := store.StringStore[key][field]
	if ok {
		t.Errorf("HRemoveStringField failed: Field '%v' still exists in StringStore with value '%v'", field, val)
	}
}

func TestHRemoveListField(t *testing.T) {
	store := storage.NewStorage()
	defer storage.ClearStorage(store)

	key := "user:1234"
	field := "genres"
	value := []string{"ITALIAN", "AMERICAN", "JAPANESE"}

	store.HSetList(key, field, value)

	store.HRemoveListField(key, field)

	val, ok := store.ListStore[key][field]
	if ok {
		t.Errorf("HRemoveListField failed: Field '%v' still exists in ListStore with values '%v'", field, val)
	}
}
