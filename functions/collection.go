package functions

import (
	"github.com/iVitaliya/colors-go"
	"github.com/iVitaliya/logger-go"
	"github.com/iVitaliya/memory-go/errors"
)

type IHasAll struct {
	// Returns if all the keys exist.
	AllExist bool
	// Returns the keys that don't exist.
	InvalidKeys []string
	// Returns the keys that do exist.
	ValidKeys []string
}

type IHasAny struct {
	// Returns whether one or several of the provided keys exist.
	HasAny bool
	// Returns the keys that don't exist.
	InvalidKeys []string
	// Returns the keys that do exist.
	ValidKeys []string
}

func KeyAt[T any](storage map[string]T, idx int) string {
	var (
		value   string = ""
		lastIdx        = len(storage) - 1
	)

	if idx > lastIdx {
		logger.Error(FormatString("The provided index \"{0}\" isn't a valid index as it exceeds the last available index", []string{
			colors.Red(Stringify[int](idx)),
		}))

		return value
	}

	for i, v := range SendAll(storage).Keys {
		if idx == i {
			value = v
		}
	}

	return value
}

func ValueAt[T any](storage map[string]T, idx int) T {
	var (
		key   string = KeyAt[T](storage, idx)
		value        = GetItem[T](storage, key)
	)

	return value
}

func FirstItem[T any](storage map[string]T) T {
	return ValueAt[T](storage, 0)
}

func FirstKey[T any](storage map[string]T) string {
	return KeyAt[T](storage, 0)
}

func HasAll[T any](storage map[string]T, keys []string) IHasAll {
	var (
		allExist    bool     = false
		invalidKeys []string = []string{}
		validKeys   []string = []string{}
	)

	for _, k := range keys {
		for _, s := range SendAll[T](storage).Keys {
			if k != s {
				AppendItem[string](invalidKeys, k)
			}

			AppendItem[string](validKeys, k)
		}
	}

	if len(invalidKeys) == 0 {
		allExist = true
	}

	return IHasAll{
		AllExist:    allExist,
		InvalidKeys: invalidKeys,
		ValidKeys:   validKeys,
	}
}

func HasAny[T any](storage map[string]T, keys []string) IHasAny {
	var (
		anyExist    bool     = true
		invalidKeys []string = []string{}
		validKeys   []string = []string{}
	)

	for _, k := range keys {
		for _, s := range SendAll[T](storage).Keys {
			if k != s {
				AppendItem[string](invalidKeys, k)
			}

			AppendItem[string](validKeys, k)
		}
	}

	if len(validKeys) < 1 {
		anyExist = false
	}

	return IHasAny{
		HasAny:      anyExist,
		InvalidKeys: invalidKeys,
		ValidKeys:   validKeys,
	}
}

func Join[T any](storage map[string]T, devider string, seperator string, include_prefix bool) string {
	if errors.CheckIfBadDevider(devider) {
		return ""
	}

	if errors.CheckIfBadSeperator(seperator) {
		return ""
	}

	var (
		str string
		all = SendAll[T](storage)
	)
	for i := 0; i < all.Size; i++ {
		if include_prefix {
			str = str + "Key: " + all.Keys[i] + " " + devider + " Value: " + Stringify[T](all.Values[i]) + seperator
		}

		str = str + all.Keys[i] + " " + devider + " " + Stringify[T](all.Values[i]) + seperator
	}

	return str
}

func JoinKeys[T any](storage map[string]T, seperator string) string {
	if errors.CheckIfBadSeperator(seperator) {
		return ""
	}

	var str string
	for _, i := range SendAll[T](storage).Keys {
		str = str + i + seperator
	}

	return str
}

func JoinValues[T any](storage map[string]T, seperator string) string {
	if errors.CheckIfBadSeperator(seperator) {
		return ""
	}

	var str string
	for _, i := range SendAll[T](storage).Values {
		str = str + Stringify[T](i) + seperator
	}

	return str
}

func LastItem[T any](storage map[string]T) T {
	var idx = len(storage) - 1

	return ValueAt[T](storage, idx)
}

func LastKey[T any](storage map[string]T) string {
	var idx = len(storage) - 1

	return KeyAt[T](storage, idx)
}
