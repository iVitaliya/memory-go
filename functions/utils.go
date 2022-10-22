package functions

import (
	"fmt"
	"strings"
)

func AppendItem[T any](array []T, item T) []T {
	var arr []T = append(array, item)

	return arr
}

func Stringify[T any](value T) string {
	return fmt.Sprint(value)
}

func FormatString(str string, value []string) string {
	formatable := func(number int) string { return "{" + fmt.Sprint(number) + "}" }

	for i := 0; i < len(value); i++ {
		str = strings.Replace(str, formatable(i), value[i], -1)
	}

	return str
}
