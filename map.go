package memory

import (
	"github.com/iVitaliya/memory-go/functions"
)

type IMap[T any] struct {
	// Returns all the existing keys in the Map.
	All func() functions.IMapAll[T]
	// The Clear() method removes all elements from the Map. Returns whether the map has been cleared or not.
	Clear func() bool
	// Removes any value associated to the key. Map().Has(key) will return false afterwards.
	Delete func(key string) bool
	// Returns a Boolean asserting whether a value has been associated to the key in the Map or not.
	Exists func(key string) bool
	// Returns the value associated to the key, or nil if there is none.
	Fetch func(key string) T
	// Runs a test over every item in the Map.
	ForEach func(function func(key string, value T) any)
	// Returns the value associated to the key, or nil if there is none. NOTE: for stricter fetching you may want to consider using Map().fetch(key)
	Get func(key string) T
	// Returns a Boolean asserting whether a value has been associated to the key in the Map or not.
	Has func(key string) bool
	// The iterator which is used to iterate through the map.
	Iterator func(function func(key string, value T) any)
	// Returns all the keys as an array.
	Keys func() []string
	// Sets the value for the key in the Map. Returns whether the key value has been set.
	Set func(key string, value T) bool
	// Returns the amount of keys in the Map.
	Size func() int
	// Updates the value for the provided key in the Map. Returns whether the key value has been changed.
	Update func(key string, value T) bool
	// Returns all the values as an array.
	Values func() []T
}

// Creates a new Map instance.
func Map[T any]() *IMap[T] {
	var store = make(map[string]T)

	var Funcs *IMap[T] = &IMap[T]{
		All: func() functions.IMapAll[T] {
			return functions.SendAll[T](store)
		},
		Clear: func() bool {
			return functions.ClearMap[T](store)
		},
		Delete: func(key string) bool {
			return functions.DeleteKey[T](store, key)
		},
		Exists: func(key string) bool {
			return functions.Exists[T](store, key)
		},
		Fetch: func(key string) T {
			return functions.FetchItem[T](store, key)
		},
		ForEach: func(function func(key string, value T) any) {
			for k, v := range store {
				function(k, v)
			}
		},
		Get: func(key string) T {
			return functions.GetItem[T](store, key)
		},
		Has: func(key string) bool {
			return functions.Exists[T](store, key)
		},
		Iterator: func(function func(key string, value T) any) {
			functions.Iterator[T](store, function)
		},
		Keys: func() []string {
			return functions.KeysSender[T](store)
		},
		Set: func(key string, value T) bool {
			return functions.Set[T](store, key, value)
		},
		Size: func() int {
			return functions.Length[T](store)
		},
		Update: func(key string, value T) bool {
			return functions.Update[T](store, key, value)
		},
		Values: func() []T {
			return functions.ValuesSender[T](store)
		},
	}

	return Funcs
}
