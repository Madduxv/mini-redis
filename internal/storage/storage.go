package storage

type Storage struct {
  StringStore map[string]map[string]string
  ListStore map[string]map[string][]string
}

func NewStorage() *Storage {
  return &Storage{
    StringStore: make(map[string]map[string]string),
    ListStore: make(map[string]map[string][]string),
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

func (s *Storage) LPush(key string, field string, value string) {

}

func (s *Storage) RPush(key string, field string, value string) {

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

