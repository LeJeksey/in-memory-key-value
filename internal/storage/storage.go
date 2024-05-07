package storage

type storage struct {
	hash map[string]string
}

func NewStorage() *storage {
	return &storage{
		hash: make(map[string]string),
	}
}

func (s *storage) Set(key, value string) {
	s.hash[key] = value
}

func (s *storage) Get(key string) string {
	return s.hash[key]
}

func (s *storage) Delete(key string) {
	delete(s.hash, key)
}
