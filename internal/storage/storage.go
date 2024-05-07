package storage

type Storage struct {
	hash map[string]string
}

func NewStorage() *Storage {
	return &Storage{
		hash: make(map[string]string),
	}
}

func (s *Storage) Set(key, value string) error {
	s.hash[key] = value
	return nil
}

func (s *Storage) Get(key string) (string, error) {
	return s.hash[key], nil
}

func (s *Storage) Delete(key string) error {
	delete(s.hash, key)
	return nil
}
