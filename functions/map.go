package functions

type iMap[T any] struct{
	items map[string]iMapItem[T]
	amount int
}


type iMapItem[T any] struct {
	Key string
	Value T
}

func CreateMap[T any]() *iMap[T] {
	mapItems := make(map[string]iMapItem[T])

	return &iMap[T]{
		items: mapItems,
		amount: 0,
	}
}

func 