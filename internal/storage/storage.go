package storage

type Storage struct {
  storeString map[string]map[string]string
  storeList map[string]map[string][]string
}

func NewStorage() *Storage {
  return &Storage{
    storeString: make(map[string]map[string]string),
    storeList: make(map[string]map[string][]string),
  }
}

func (s *Storage) HSet(key string, field string, value string) {
    // Check if the key already exists in the store
    if _, exists := s.storeString[key]; !exists {
        // If not, create a new hash map for this key
        s.storeString[key] = make(map[string]string)
    }

    // Set the field in the hash
    s.storeString[key][field] = value
}

func (s *Storage) HGet(key string, field string) (string, bool) {
    if fields, exists := s.storeString[key]; exists {
        if value, fieldExists := fields[field]; fieldExists {
            return value, true
        }
    }
    return "", false
}
