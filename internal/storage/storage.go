package storage

import "errors"

type Storage struct {
	HashStore map[string]map[string]string
	// ListStore       map[string]map[string][]string
	LinkedListStore map[string]map[string]*LinkedList
	SetStore        map[string][]string
}

type Node struct {
	Data string
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func NewStorage() *Storage {
	return &Storage{
		HashStore: make(map[string]map[string]string),
		// ListStore:       make(map[string]map[string][]string),
		LinkedListStore: make(map[string]map[string]*LinkedList),
		SetStore:        make(map[string][]string),
	}
}

func ClearStorage(s *Storage) {
	for key := range s.HashStore {
		delete(s.HashStore, key)
	}
	// for key := range s.ListStore {
	// 	delete(s.ListStore, key)
	// }
	for key := range s.LinkedListStore {
		delete(s.LinkedListStore, key)
	}
	for key := range s.SetStore {
		delete(s.SetStore, key)
	}
}

func (s *Storage) Del(key string) {
	delete(s.HashStore, key)
	delete(s.SetStore, key)
	delete(s.LinkedListStore, key)
}

func (s *Storage) HSet(key, field, value string) {
	if _, hashStoreExists := s.HashStore[key]; !hashStoreExists {
		s.HashStore[key] = make(map[string]string)
	}
	s.HashStore[key][field] = value
}

func (s *Storage) HGet(key, field string) (string, bool) {
	if fields, exists := s.HashStore[key]; exists {
		if value, fieldExists := fields[field]; fieldExists {
			return value, true
		}
	}
	return "", false
}

func (s *Storage) HRem(key string) {
	if _, exists := s.HashStore[key]; exists {
		delete(s.HashStore, key)
	}
}

func (s *Storage) HDel(key, field string) bool {
	if _, exists := s.HashStore[key]; exists {
		delete(s.HashStore[key], field)
		return true
	}
	return false
}

func (s *Storage) SAdd(key, value string) {
	if _, exists := s.SetStore[key]; !exists {
		s.SetStore[key] = make([]string, 0)
	}
	s.SetStore[key] = append(s.SetStore[key], value)
}

func (s *Storage) SGet(key string) ([]string, bool) {
	if _, exists := s.SetStore[key]; !exists {
		return nil, false
	}
	return s.SetStore[key], true
}

func (s *Storage) SRem(key, value string) int8 {
	if _, exists := s.SetStore[key]; !exists {
		return 0
	}
	if len(s.SetStore[key]) == 0 {
		return 0
	}
	for i := 0; i < len(s.SetStore[key]); i++ {
		if s.SetStore[key][i] == value {
			s.SetStore[key] = append(s.SetStore[key][:i], s.SetStore[key][i+1:]...)
			return 1
		}
	}
	return 0
}

func (s *Storage) RPush(key, field, value string) {
	if _, keyExists := s.LinkedListStore[key]; !keyExists {
		s.LinkedListStore[key] = make(map[string]*LinkedList)
	}

	if _, fieldExists := s.LinkedListStore[key][field]; !fieldExists {
		s.LinkedListStore[key][field] = &LinkedList{}
	}

	list := s.LinkedListStore[key][field]
	newNode := &Node{Data: value, Next: nil}

	if list.Head == nil {
		list.Head = newNode
		return
	}

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (s *Storage) LRange(key, field string, start, end int) []string {
	list, keyExists := s.LinkedListStore[key][field]
	if !keyExists || list.Head == nil {
		return nil
	}

	length := 0
	current := list.Head
	for current != nil {
		length++
		current = current.Next
	}

	if start < 0 {
		start = length + start
	}
	if end < 0 {
		end = length + end
	}

	if start < 0 {
		start = 0
	}
	if end >= length {
		end = length - 1
	}
	if start > end {
		return nil
	}

	result := []string{}
	current = list.Head
	index := 0

	for current != nil && index <= end {
		if index >= start {
			result = append(result, current.Data)
		}
		current = current.Next
		index++
	}

	return result
}

func (s *Storage) LClear(key string, field string) error {
	listMap, keyExists := s.LinkedListStore[key]
	if !keyExists {
		return errors.New("key does not exist")
	}

	_, fieldExists := listMap[field]
	if !fieldExists {
		return errors.New("field does not exist")
	}

	// Clear the list by setting the head to nil
	listMap[field] = &LinkedList{Head: nil}
	return nil
}

// func (s *Storage) HSetList(key, field string, value []string) {
// 	if _, ListStoreExists := s.ListStore[key]; !ListStoreExists {
// 		s.ListStore[key] = make(map[string][]string)
// 	}
// 	s.ListStore[key][field] = value
// }

// func (s *Storage) HGetList(key, field string) ([]string, bool) {
// 	if fields, exists := s.ListStore[key]; exists {
// 		if value, fieldExists := fields[field]; fieldExists {
// 			return value, true
// 		}
// 	}
// 	return nil, false
// }

// func (s *Storage) HRemoveListField(key, field string) bool {
// 	if _, exists := s.ListStore[key]; exists {
// 		delete(s.ListStore[key], field)
// 		return true
// 	}
// 	return false
// }

//	func (s *Storage) HAdd(key, field, value string) {
//		if _, exists := s.ListStore[key]; !exists {
//			s.ListStore[key] = make(map[string][]string)
//			s.ListStore[key][field] = make([]string, 0)
//		}
//		s.ListStore[key][field] = append(s.ListStore[key][field], value)
//	}

// func (s *Storage) HRem(key, field, value string) int8 {
// 	if _, exists := s.ListStore[key][value]; !exists {
// 		return 0
// 	}
// 	if len(s.ListStore[key][value]) == 0 {
// 		return 0
// 	}
// 	for i := 0; i < len(s.ListStore[key][field]); i++ {
// 		if s.ListStore[key][field][i] == value {
// 			s.ListStore[key][field] = append(s.ListStore[key][field][:i], s.ListStore[key][field][i+1:]...)
// 			return 1
// 		}
// 	}
// 	return 0
// }
