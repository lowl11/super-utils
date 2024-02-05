package validators

import (
	"regexp"
	"strings"
)

// availableSymbols Допустимые символы только: буквы латиницы и кириллицы, пробелы, дефисы/тире, точки
var availableSymbols = regexp.MustCompile(`^[A-Za-zА-Яа-яҚқҢңӘәҰұҮүІіЁёҒғөӨһҺ\s\-—\.]+$`)
var cyrillicWithKazSymbolsAndCymbol = regexp.MustCompile(`^[А-Яа-яҚқҢңӘәҰұҮүІіЁёҒғөӨһҺ\s\-—\.]+$`)
var latinSymbols = regexp.MustCompile(`^[A-Za-z\s\-—\.]+$`)

// Создаем регулярное выражение для поиска последовательных пробелов длиной больше одного
var more2Space = regexp.MustCompile(`\s{2,}`)

// ValidateFirstName alfa version
// if data is empty return false
func ValidateFirstName(data string) (string, bool) {
	data = clearingSpace(data)

	//check if data is empty
	if len(data) == 0 {
		return data, false
	}

	return basicValidationForFio(data)
}

// ValidateLastName alfa version
// if data is empty return false
func ValidateLastName(data string) (string, bool) {
	data = clearingSpace(data)

	//check if data is empty
	if len(data) == 0 {
		return data, false
	}

	return basicValidationForFio(data)
}

// ValidateMiddleName alfa version
// if data is empty return true
func ValidateMiddleName(data string) (string, bool) {
	data = clearingSpace(data)

	//check if data is empty
	if len(data) == 0 {
		return data, true
	}

	return basicValidationForFio(data)
}

// basicValidationForFio
// if data is empty return false
func basicValidationForFio(data string) (string, bool) {
	//check if data contains only cyrillic symbols
	ok := availableSymbols.MatchString(data)
	if !ok {
		return data, false
	}

	//check if data contains only cyrillic symbols
	ok = cyrillicWithKazSymbolsAndCymbol.MatchString(data)
	if !ok {
		//check if data contains only latin symbols
		ok = latinSymbols.MatchString(data)
		if !ok {
			return data, false
		}
	}

	return data, true
}
func clearingSpace(data string) string {
	data = strings.TrimLeft(data, " ")
	data = strings.TrimRight(data, " ")
	data = more2Space.ReplaceAllString(data, " ")
	return data
}
