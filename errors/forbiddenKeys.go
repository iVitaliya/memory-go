package errors

import (
	"strings"

	"github.com/iVitaliya/colors-go"
	"github.com/iVitaliya/logger-go"
	"github.com/iVitaliya/logger-go/utils"
)

// =====================
//
// # INVALID CHARACTERS TO USE
//
// =====================
var (
	invalidCharacters = []string{
		"?",
		"!",
		"\\",
		"/",
		"|",
		"\"",
		"'",
		":",
		";",
		"[",
		"]",
		"{",
		"}",
		"=",
		"\u200b",
		"@",
		"#",
		"$",
		"%",
		"^",
		"&",
		"*",
		"(",
		")",
	}

	// ============
	//
	// # BANNED TO USE
	//
	// ============
	bannedKeys = []string{
		"pussy",
		"p*ssy",
		"p_ssy",
		"p**sy",
		"p__sy",
		"dick",
		"d*ck",
		"d_ck",
		"cock",
		"cuck",
		"c*ck",
		"c_ck",
		"c**k",
		"c__k",
		"penis",
		"p*nis",
		"p_nis",
		"p**is",
		"p__is",
		"fuck",
		"whore",
		"hoe",
		"cunt",
		"f*ck",
		"f**k",
		"f***",
	}

	// ==============
	//
	// # PROHIBITED TO USE
	//
	// ==============
	badSeperators = []string{
		"\u200b",
		"!",
		"?",
		"+",
		"=",
		"@",
		"#",
		"$",
		"%",
		"^",
		"&",
		"*",
		"(",
		")",
		"_",
		"`",
		"\"",
		"'",
		";",
	}
	badDeviders = []string{
		"\u200b",
		"!",
		"?",
		"+",
		")",
		"(",
		"^",
		"%",
		"$",
		"`",
		"\"",
		"'",
		";",
	}
)

func CheckIfBadKey(key string) bool {
	for _, v := range invalidCharacters {
		str := strings.Contains(key, v)
		if str {
			logger.Error(utils.FormatString("You defined \"{0}\" in your key as seperator which isn't an allowed character, please use one of the following characters instead as seperator: \"-\", \".\", \"_\"", []string{
				colors.Red(v),
			}))

			return true
		}
	}

	for _, v := range bannedKeys {
		str := strings.Contains(key, v)
		if str {
			logger.Error(utils.FormatString("You defined \"{0}\" in your key which isn't an allowed word, please refractor your key!", []string{
				colors.Red(v),
			}))

			return true
		}
	}

	return false
}

func CheckIfBadSeperator(value string) bool {
	for _, v := range badSeperators {
		if v == value {
			logger.LogFormatted(logger.ErrorState, "The character \"{0}\" isn't allowed to be used as a joinable seperator, please use one of the following characters instead as seperator: \"-\", \".\", \",\", \"|\"", []string{
				colors.Red(v),
			})

			return true
		}
	}

	return false
}

func CheckIfBadDevider(value string) bool {
	for _, v := range badDeviders {
		if v == value {
			logger.LogFormatted(logger.ErrorState, "The character \"{0}\" isn't allowed to be used as a joinable devider, please use one of the following characters instead as devider: \"-\", \",\", \"|\", \":\", \"@\", \"#\", \"~\"", []string{
				colors.Red(v),
			})

			return true
		}
	}

	return false
}