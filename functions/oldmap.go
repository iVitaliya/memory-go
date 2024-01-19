package functions

import (
	"github.com/iVitaliya/memory-go/errors"
)

func Exists[T any](storage map[string]T, key string) bool {
	for k, _ := range storage {
		if k == key {
			return true
		}
	}

	return false
}

func Length[T any](storage map[string]T) int {
	length := len(storage)
	if length == -1 || length < 1 {
		return 0
	}

	return length
}

func Set[T any](storage map[string]T, key string, value T) bool {
	if errors.CheckIfBadKey(key) {
		return false
	}

	if key == "" {
		errors.NoKeyError(errors.Set, key)
		return false
	}

	storage[key] = value
	return true
}

func Update[T any](storage map[string]T, key string, value T) bool {
	if !Exists[T](storage, key) {
		errors.NoKeyError(errors.Update, key)

		return false
	}

	return Set[T](storage, key, value)
}

type IMapAll[T any] struct {
	Keys   []string
	Values []T
	Size   int
}

func SendAll[T any](storage map[string]T) IMapAll[T] {
	if len(storage) < 1 {
		errors.NoKeysError(errors.All, 2)

		return IMapAll[T]{}
	}

	var (
		res    IMapAll[T]
		keys   []string
		values []T
		size   int
	)

	for k, v := range storage {
		keys = AppendItem[string](keys, k)
		values = AppendItem[T](values, v)
		size = size + 1
	}

	res = IMapAll[T]{
		Keys:   keys,
		Values: values,
		Size:   size,
	}

	return res
}

func FetchItem[T any](storage map[string]T, key string) T {
	res, ok := storage[key]
	if ok {
		return res
	}

	errors.KeyNotFoundError(errors.Get, key)
	return *new(T)
}

func GetItem[T any](storage map[string]T, key string) T {
	for k, v := range storage {
		if k == key {
			return v
		}
	}

	errors.KeyNotFoundError(errors.Get, key)
	return *new(T)
}

func DeleteKey[T any](storage map[string]T, key string) bool {
	if !Exists[T](storage, key) {
		errors.NoKeyError(errors.Delete, key)

		return false
	}

	delete(storage, key)
	return true
}

func ClearMap[T any](storage map[string]T) bool {
	if len(storage) < 1 {
		errors.NoKeysError(errors.Clear, 2)

		return false
	}

	for k, _ := range storage {
		delete(storage, k)
	}

	return true
}

func KeysSender[T any](storage map[string]T) []string {
	if len(storage) < 1 {
		errors.NoKeysError(errors.Keys, 2)

		return []string{}
	}

	var res []string
	for k, _ := range storage {
		res = append(res, k)
	}

	return res
}

func ValuesSender[T any](storage map[string]T) []T {
	if len(storage) < 1 {
		errors.NoKeysError(errors.Values, 2)

		return []T{}
	}

	var res []T
	for _, v := range storage {
		res = append(res, v)
	}

	return res
}

func Iterator[T any](storage map[string]T, function func(key string, value T) any) {
	for k, v := range storage {
		function(k, v)
	}
}
