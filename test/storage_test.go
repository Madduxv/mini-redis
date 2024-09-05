package test

import (
	"github.com/Madduxv/mini-redis/internal/storage"
	"testing"
)

func TestHSet(t *testing.T) {
	store := storage.NewStorage()
	defer storage.ClearStorage(store)

	key := "user:1234"
	field := "name"
	value := "Maddux"

	store.HSet(key, field, value)

	if store.HashStore[key][field] != value {
		t.Errorf("HSet failed, expected %v, got %v", value, store.HashStore[key][field])
	}
}

func TestSetGetRem(t *testing.T) {
	store := storage.NewStorage()
	defer storage.ClearStorage(store)

	var key string = "key:1234"
	var value string = "value:1234"

	if noVal := store.Get(key); noVal != "" {
		t.Errorf("Get got data before anything was added: Found %s", noVal)
	}

	store.Set(key, value)

	if val := store.Store[key]; val != value {
		t.Errorf("Set failed: Value was not added correctly to store. Found %s", val)
	}
	if recievedVal := store.Get(key); recievedVal != store.Store[key] {
		t.Errorf("Get Failed: Expected %s, but found %s", store.Store[key], recievedVal)
	}

	store.Rem(key)

	if val, exists := store.Store[key]; exists {
		t.Errorf("Rem Failed: Found %s", val)
	}

}

func TestRPushLRange(t *testing.T) {
	store := storage.NewStorage()
	defer storage.ClearStorage(store)

	key := "user:1234"
	field := "genres"
	value0 := "ITALIAN"
	value1 := "AMERICAN"
	value2 := "JAPANESE"

	store.RPush(key, field, value0)
	store.RPush(key, field, value1)
	store.RPush(key, field, value2)

	data := store.LRange(key, field, 0, -1)

	if data == nil {
		t.Error("Data either not going in or not coming out for RPush and LRange")
	}
}
func TestLClear(t *testing.T) {
	store := storage.NewStorage()

	// Set up an initial linked list with some data
	key := "mylist"
	field := "field1"
	store.LinkedListStore[key] = map[string]*storage.LinkedList{
		field: {Head: &storage.Node{Data: "first", Next: &storage.Node{Data: "second", Next: nil}}},
	}

	// Call LClear to clear the list
	err := store.LClear(key, field)
	if err != nil {
		t.Fatalf("LClear returned an error: %v", err)
	}

	// Check if the head of the list is nil
	if store.LinkedListStore[key][field].Head != nil {
		t.Error("Data not being removed from LinkedListStore")
	}
}

func TestSAddGetRem(t *testing.T) {
	store := storage.NewStorage()
	defer storage.ClearStorage(store)

	key := "Maddux's Group"
	value := "Maddux"
	value1 := "Trin"

	store.SAdd(key, value)
	store.SAdd(key, value1)

	if len(store.SetStore[key]) != 2 {
		t.Errorf("SAdd failed: Items were not added to the list (Found: %v)", store.SetStore[key])
	}

	store.SRem(key, value1)

	if len(store.SetStore[key]) != 1 {
		t.Errorf("Rem failed: Items were not removed from the list (Found: %v)", store.SetStore[key])
	}

	if returned_value, exists := store.SGet(key); !exists {
		t.Error("SGet says that the field does not exist")
	} else {
		t.Logf("SGet: expected: %v, returned %v", returned_value, store.SetStore[key])
	}

}

func TestHRem(t *testing.T) {
	store := storage.NewStorage()
	defer storage.ClearStorage(store)

	key := "user:1234"
	field := "name"
	value := "Maddux"

	store.HSet(key, field, value)
	store.HRem(key)

	_, ok := store.HashStore[key]
	if ok {
		t.Errorf("HRemove failed: Key '%v' still exists in StringStore", key)
	}

}

func TestHDel(t *testing.T) {
	store := storage.NewStorage()
	defer storage.ClearStorage(store)

	key := "user:1234"
	field := "name"
	value := "Maddux"

	store.HSet(key, field, value)

	store.HDel(key, field)

	val, ok := store.HashStore[key][field]
	if ok {
		t.Errorf("HDel failed: Field '%v' still exists in StringStore with value '%v'", field, val)
	}
}

func TestDel(t *testing.T) {
	store := storage.NewStorage()
	defer storage.ClearStorage(store)

	key := "user:1234"
	field := "name"
	value := "Maddux"

	store.HSet(key, field, value)

	field1 := "genres"
	value1 := "ITALIAN"
	value2 := "JAPANESE"

	store.RPush(key, field1, value1)
	store.RPush(key, field1, value2)

	store.Del(key)

	if hash, hashExists := store.HashStore[key]; hashExists {
		t.Errorf("Key was not removed from HashStore: Found %v", hash)
	}
	if node, linkedListExists := store.LinkedListStore[key]; linkedListExists {
		t.Errorf("Key was not removed from LinkedListStore: Found %v", node)
	}

}
