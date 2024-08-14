package test

import "testing"
import "github.com/Madduxv/mini-redis/internal/storage"

func TestHSet(t *testing.T) {
    storage := storage.NewStorage()
    
    key := "user:1234"
    field := "name"
    value := "Maddux"

    storage.HSet(key, field, value)

    if storage.StringStore[key][field] != value {
        t.Errorf("HSet failed, expected %v, got %v", value, storage.StringStore[key][field])
    }
}

func TestHSetList(t *testing.T) {
    storage := storage.NewStorage()
    
    key := "user:1234"
    field := "genres"
    value := []string{"ITALIAN", "AMERICAN", "JAPANESE"}

    storage.HSetList(key, field, value)

    if len(storage.ListStore[key][field]) != len(value) {
        t.Errorf("HSetList failed, expected list of length %d, got %d", len(value), len(storage.ListStore[key][field]))
    }

    for i, v := range value {
        if storage.ListStore[key][field][i] != v {
            t.Errorf("HSetList failed, expected %v at index %d, got %v", v, i, storage.ListStore[key][field][i])
        }
    }
}
