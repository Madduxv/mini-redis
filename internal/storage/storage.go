package storage

type Storage struct {
  stringStore map[string]map[string]string
  listStore map[string]map[string][]string
}

func NewStorage() *Storage {
  return &Storage{
    stringStore: make(map[string]map[string]string),
    listStore: make(map[string]map[string][]string),
  }
}

func (s *Storage) HSet(key string, field string, value string) {
  if _, exists := s.stringStore[key]; !exists {
    s.stringStore[key] = make(map[string]string)
  }
  s.stringStore[key][field] = value
}

func (s *Storage) HGet(key string, field string) (string, bool) {
  if fields, exists := s.stringStore[key]; exists {
    if value, fieldExists := fields[field]; fieldExists {
      return value, true
    }
  }
  return "", false
}

func (s *Storage) HRemove(key string) {
  if _, exists := s.stringStore[key]; exists {
    delete(s.stringStore, key)
  }
  if _, exists := s.listStore[key]; exists {
    delete(s.listStore, key)
  }
}

func (s *Storage) HRemoveStringField(key, field string) bool {
  if _, exists := s.stringStore[key]; exists {
    delete(s.stringStore[key], field)
    return true
  }
  return false
}

func (s *Storage) HRemoveListField(key, field string) bool {
  if _, exists := s.stringStore[key]; exists {
    delete(s.listStore[key], field)
    return true
  }
  return false
}

func (s *Storage) HSetList(key string, field string, value []string) {
  if _, exists := s.stringStore[key]; !exists {
    s.listStore[key] = make(map[string][]string)
  }
  s.listStore[key][field] = value
} 

func (s *Storage) HGetList(key string, field string) ([]string, bool) {
  if fields, exists := s.listStore[key]; exists {
    if value, fieldExists := fields[field]; fieldExists {
      return value, true
    }
  }
  return nil, false
}
