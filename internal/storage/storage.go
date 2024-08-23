package storage

type Storage struct {
	StringStore     map[string]map[string]string
	ListStore       map[string]map[string][]string
	LinkedListStore map[string]map[string]*LinkedList
}

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewStorage() *Storage {
	return &Storage{
		StringStore:     make(map[string]map[string]string),
		ListStore:       make(map[string]map[string][]string),
		LinkedListStore: make(map[string]map[string]*LinkedList),
	}
}

func ClearStorage(s *Storage) {
	for key := range s.StringStore {
		delete(s.StringStore, key)
	}
	for key := range s.ListStore {
		delete(s.ListStore, key)
	}
}

func (s *Storage) HSet(key string, field string, value string) {
	if _, strStoreExists := s.StringStore[key]; !strStoreExists {
		s.StringStore[key] = make(map[string]string)
	}
	s.StringStore[key][field] = value
}

func (s *Storage) RPush(key string, field string, value string) {
	if _, keyExists := s.LinkedListStore[key]; !keyExists {
		s.LinkedListStore[key] = make(map[string]*LinkedList)
	}

	if _, fieldExists := s.LinkedListStore[key][field]; !fieldExists {
		s.LinkedListStore[key][field] = &LinkedList{}
	}

	list := s.LinkedListStore[key][field]
	newNode := &Node{data: value, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (s *Storage) LRange(key string, field string, start int, end int) []string {
	// Retrieve the linked list for the given key and field
	list, keyExists := s.LinkedListStore[key][field]
	if !keyExists || list.head == nil {
		return nil // Return nil if the key/field doesn't exist or the list is empty
	}

	// Step 1: Calculate the length of the list
	length := 0
	current := list.head
	for current != nil {
		length++
		current = current.next
	}

	// Step 2: Normalize start and end indices
	if start < 0 {
		start = length + start
	}
	if end < 0 {
		end = length + end
	}

	// Adjust indices if they are out of bounds
	if start < 0 {
		start = 0
	}
	if end >= length {
		end = length - 1
	}
	if start > end {
		return nil // Return an empty slice if start is greater than end
	}

	// Step 3: Traverse the list and collect values in the specified range
	result := []string{}
	current = list.head
	index := 0

	for current != nil && index <= end {
		if index >= start {
			result = append(result, current.data)
		}
		current = current.next
		index++
	}

	return result
}

func (s *Storage) HSetList(key string, field string, value []string) {
	if _, ListStoreExists := s.ListStore[key]; !ListStoreExists {
		s.ListStore[key] = make(map[string][]string)
	}
	s.ListStore[key][field] = value
}

func (s *Storage) HGet(key string, field string) (string, bool) {
	if fields, exists := s.StringStore[key]; exists {
		if value, fieldExists := fields[field]; fieldExists {
			return value, true
		}
	}
	return "", false
}

func (s *Storage) HGetList(key string, field string) ([]string, bool) {
	if fields, exists := s.ListStore[key]; exists {
		if value, fieldExists := fields[field]; fieldExists {
			return value, true
		}
	}
	return nil, false
}

func (s *Storage) HRemove(key string) {
	if _, exists := s.StringStore[key]; exists {
		delete(s.StringStore, key)
	}
	if _, exists := s.ListStore[key]; exists {
		delete(s.ListStore, key)
	}
}

func (s *Storage) HRemoveStringField(key, field string) bool {
	if _, exists := s.StringStore[key]; exists {
		delete(s.StringStore[key], field)
		return true
	}
	return false
}

func (s *Storage) HRemoveListField(key, field string) bool {
	if _, exists := s.ListStore[key]; exists {
		delete(s.ListStore[key], field)
		return true
	}
	return false
}
