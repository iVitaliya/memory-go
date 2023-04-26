package errors

import (
	"github.com/iVitaliya/colors-go"
	"github.com/iVitaliya/logger-go"
	"github.com/iVitaliya/logger-go/utils"
)

const (
	Delete uint = 1
	Clear  uint = 2
	Set    uint = 3
	Update uint = 4
	Get    uint = 5
	All    uint = 6

	Keys   = 7
	Values = 8
)

func NoKeysError(msgType uint, logEmptySpaces int) {
	logger.LogEmptySpace(logEmptySpaces)

	var str string
	switch msgType {
	case Clear:
		str = "There are no keys set in the map and thus there are no keys to clear from the map"
		break
	case All:
		str = "There are no keys set in the map and thus I can't display any keys/values"
		break
	case Keys:
		str = "There are no keys set in the map and thus I can't display any keys"
		break
	case Values:
		str = "There are no keys set in the map and thus I can't display any values"
		break
	}

	logger.Error(str)
}

func NoKeyError(msgType uint, key string) {
	logger.LogEmptySpace(2)

	var str string
	switch msgType {
	case Delete:
		str = utils.FormatString("The key \"{0}\" couldn't be deleted as it either doesn't exist or it failed deleting it", []string{
			colors.Red(key),
		})
		break
	case Set:
		str = utils.FormatString("The key \"{0}\" couldn't be set in the map as the provided key was either empty or it contained invalid characters", []string{
			colors.Red(key),
		})
		break
	case Update:
		str = utils.FormatString("The key \"{0}\" couldn't be updated in the map as it couldn't be found as a existing key", []string{
			colors.Red(key),
		})
		break
	}

	logger.Error(str)
}

func KeyNotFoundError(msgType uint, key string) {
	logger.LogEmptySpace(2)

	var str string
	switch msgType {
	case Get:
		str = utils.FormatString("The key \"{0}\" couldn't be found in the map, are you sure you spelled it correctly?", []string{
			colors.Dim(colors.BrightGreen(key)),
		})
		break
	}

	logger.Error(str)
}
