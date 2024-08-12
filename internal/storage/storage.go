package storage

type Storage struct {
  store map[string]map[string]string
}

func NewStorage() *Storage {
  return &Storage{
    store: make(map[string]map[string]string),
  }
}

// HSet sets a field in the hash stored at key to value.
func (s *Storage) HSet(key string, field string, value string) {
    // Check if the key already exists in the store
    if _, exists := s.store[key]; !exists {
        // If not, create a new hash map for this key
        s.store[key] = make(map[string]string)
    }

    // Set the field in the hash
    s.store[key][field] = value
}

// HGet gets the value of a field in the hash stored at key.
func (s *Storage) HGet(key string, field string) (string, bool) {
    if fields, exists := s.store[key]; exists {
        if value, fieldExists := fields[field]; fieldExists {
            return value, true
        }
    }
    return "", false
}
