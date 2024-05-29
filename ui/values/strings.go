package values

import (
	"encoding/json"

	"github.com/topwuther/replayer/ui/values/localizable"
)

const (
	ENGLISH = iota
	CHINESE = iota
)

var (
	userLanguage = ENGLISH
	languages    = []int{ENGLISH, CHINESE}
)

func getLanguageField() []byte {
	switch userLanguage {
	case ENGLISH:
		return localizable.EN
	case CHINESE:
		return localizable.ZH
	default:
		return localizable.EN
	}
}

func hasLanguage(language int) bool {
	for _, value := range languages {
		if value == language {
			return true
		}
	}
	return false
}

func SetUserLanguage(lang int) {
	if hasLanguage(lang) {
		userLanguage = lang
	}
}

func GetText(key string) string {
	lang := getLanguageField()
	var data map[string]string

	err := json.Unmarshal(lang, &data)
	if err != nil {
		return "<nil>"
	}
	text := data[key]
	if text == "" {
		return key
	}
	return text
}
