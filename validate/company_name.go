package validate

import (
	"regexp"
)

func CompanyName(companyName string) bool {
	// Проверка на допустимые символы
	allowedChars := regexp.MustCompile(`^[a-zA-Zа-яА-Я0-9"' \-,./«»]+$`)
	if !allowedChars.MatchString(companyName) {
		return false
	}
	return true
}
