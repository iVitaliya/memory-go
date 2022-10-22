package memory

import "github.com/iVitaliya/memory-go/functions"

type ICollection[T any] struct {
	// Returns all the existing keys in the Collection.
	All func() functions.IMapAll[T]
	// Returns the item at a given index, allowing for positive and negative integers. If it fails then it'll return nil.
	At func(idx int) T
	// The Clear() method removes all elements from the Collection. Returns whether the collection has been cleared or not.
	Clear func() bool
	// Removes any value associated to the key. Collection().has(key) will return false afterwards.
	Delete func(key string) bool
	// Returns a Boolean asserting whether a value has been associated to the key in the Collection or not.
	Exists func(key string) bool
	// Returns the value associated to the key, or nil if there is none.
	Fetch func(key string) T
	// Returns the first value from this Collection. If there is none or if it fails then it'll return nil.
	First func() T
	// Returns the first associated key in the Collection. If there is none or if it fails then it'll return an empty string.
	FirstKey func() string
	// Runs the given function in a for-loop/for-each-loop.
	ForEach func(function func(key string, value T) any)
	// Returns the value associated to the key, or nil if there is none. NOTE: for stricter fetching you may want to consider using Collection().fetch(key)
	Get func(key string) T
	// Returns a Boolean asserting whether a value has been associated to the key in the Collection or not.
	Has func(key string) bool
	// Checks if all of the elements exist in the Collection.
	HasAll func(keys []string) functions.IHasAll
	// Checks if any of the provided keys exist in the Collection. If one exists it'll return if any exist and the key that exists.
	HasAny func(keys []string) functions.IHasAny
	// The iterator which is used to iterate through the map.
	Iterator func(function func(key string, value T) any)
	// Joins the keys and values by provided devider and joins the line by seperator.
	//
	// If include prefix is set to true then it wil prefix the keys or values with "Key: " or "Value: ".
	//
	// If the join process fails, it'll send an empty string.
	Join func(devider string, seperator string, include_prefix bool) string
	// Joins the keys with the given seperator.
	//
	// If the join process fails, it'll send an empty string.
	JoinKeys func(seperator string) string
	// Joins the values with the given seperator.
	//
	// If the join process fails, it'll send an empty string.
	JoinValues func(seperator string) string
	// Returns all the keys as an array.
	Keys func() []string
	// Returns the key at a given index, allowing for positive and negative integers. If it fails then it'll return an empty string.
	KeyAt func(idx int) string
	// Returns the last value from this Collection. If there is none or if it fails then it'll return nil.
	Last func() T
	// Returns the last associated key in the Collection. If there is none or if it fails then it'll return an empty string.
	LastKey func() string
	// Sets the value for the key in the Collection. Returns whether the key value has been set.
	Set func(key string, value T) bool
	// Returns the amount of keys in the Collection.
	Size func() int
	// Updates the value for the provided key in the Collection. Returns whether the key value has been changed.
	Update func(key string, value T) bool
	// Returns all the values as an array.
	Values func() []T
}

func Collection[T any]() *ICollection[T] {
	var (
		store                 = make(map[string]T)
		Funcs *ICollection[T] = &ICollection[T]{
			All: func() functions.IMapAll[T] {
				return functions.SendAll[T](store)
			},
			At: func(idx int) T {
				return functions.ValueAt[T](store, idx)
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
			First: func() T {
				return functions.FirstItem[T](store)
			},
			FirstKey: func() string {
				return functions.FirstKey[T](store)
			},
			ForEach: func(function func(key string, value T) any) {
				functions.Iterator[T](store, function)
			},
			Get: func(key string) T {
				return functions.GetItem[T](store, key)
			},
			Has: func(key string) bool {
				return functions.Exists[T](store, key)
			},
			HasAll: func(keys []string) functions.IHasAll {
				return functions.HasAll[T](store, keys)
			},
			HasAny: func(keys []string) functions.IHasAny {
				return functions.HasAny[T](store, keys)
			},
			Iterator: func(function func(key string, value T) any) {
				functions.Iterator[T](store, function)
			},
			Join: func(devider string, seperator string, include_prefix bool) string {
				return functions.Join[T](store, devider, seperator, include_prefix)
			},
			JoinKeys: func(seperator string) string {
				return functions.JoinKeys[T](store, seperator)
			},
			JoinValues: func(seperator string) string {
				return functions.JoinValues[T](store, seperator)
			},
			Keys: func() []string {
				return functions.SendAll[T](store).Keys
			},
			KeyAt: func(idx int) string {
				return functions.KeyAt[T](store, idx)
			},
			Last: func() T {
				return functions.LastItem[T](store)
			},
			LastKey: func() string {
				return functions.LastKey[T](store)
			},
			Set: func(key string, value T) bool {
				return functions.Set[T](store, key, value)
			},
			Size: func() int {
				return functions.SendAll[T](store).Size
			},
			Update: func(key string, value T) bool {
				return functions.Update[T](store, key, value)
			},
			Values: func() []T {
				return functions.SendAll[T](store).Values
			},
		}
	)

	return Funcs
}
