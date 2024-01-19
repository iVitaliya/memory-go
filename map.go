package memory

type iMapItem[Val any] struct {
	key   string
	value Val
}

type iMap[T any] struct {
	items map[string]iMapItem[T]
	count int
}

// Creates a new Map instance.
func Map[T any]() *iMap[T] {
	items := make(map[string]iMapItem[T])

	return &iMap[T]{
		items: items,
		count: 0,
	}
}

func (m *iMap[T]) All() []iMapItem[T] {

}

func (m *iMap[T]) Get(key string) T {
	for k, v := range *&m.items {

	}
}
