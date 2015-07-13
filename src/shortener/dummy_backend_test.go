package shortener

type DummyBackend struct {
	Backend
}

func (be *DummyBackend) GetUniqID() int64 {
	return uniqIDStartsAt
}

func NewDummyBackend() *DummyBackend {
	return &DummyBackend{}
}
