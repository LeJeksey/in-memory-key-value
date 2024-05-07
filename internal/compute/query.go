package compute

type Query interface {
	Run(storage Storage) (string, error)
}

type SetQuery struct {
	Key string
	Val string
}

func (q SetQuery) Run(storage Storage) (string, error) {
	return "", storage.Set(q.Key, q.Val)
}

type GetQuery struct {
	Key string
}

func (q GetQuery) Run(storage Storage) (string, error) {
	return storage.Get(q.Key)
}

type DeleteQuery struct {
	Key string
}

func (q DeleteQuery) Run(storage Storage) (string, error) {
	return "", storage.Delete(q.Key)
}
