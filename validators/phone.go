package validators

import (
	"regexp"
	"strings"
)

var availablePhoneAndFirstPlus = regexp.MustCompile(`^[+]?[0-9()\s-–]+$`)
var charsToRemove = "+()-–+ "

var KazakhstanOperatorCode = []string{"700", "701", "702", "703", "704", "705",
	"706", "707", "708", "709", "747", "750", "751", "760", "761",
	"762", "763", "764", "771", "775", "776", "777", "778", "301",
	"302", "303", "304", "305", "306", "307", "308", "309", "347",
	"350", "351", "360", "361", "362", "363", "364", "371", "375",
	"376", "377", "378"}

var oldToNewCode = map[string]string{
	"301": "701",
	"302": "702",
	"303": "703",
	"304": "704",
	"305": "705",
	"306": "706",
	"307": "707",
	"308": "708",
	"309": "709",
	"347": "747",
	"350": "750",
	"351": "751",
	"360": "760",
	"361": "761",
	"362": "762",
	"363": "763",
	"364": "764",
	"371": "771",
	"375": "775",
	"376": "776",
	"377": "777",
	"378": "778",
}

func ValidateMobilePhone(phone string) (string, bool) {
	if !IsValidPhone(phone) {
		return phone, false
	}
	trimmedPhone := trimMobilePhone(phone)

	if len(trimmedPhone) == 0 {
		return phone, false
	}

	if IsKazakhstanPhone(trimmedPhone) {
		trimmedPhone = transformKazNumber(trimmedPhone)
	}

	if IsRussianPhone(trimmedPhone) {
		trimmedPhone = transformRussianNumber(trimmedPhone)
	}

	return "+" + trimmedPhone, true
}

// IsValidPhone Проверяет чтоб были допустимые символы и первый знак +
func IsValidPhone(str string) bool {
	return availablePhoneAndFirstPlus.MatchString(str)
}

func isFirstPlus(phone string) bool {
	if len(phone) == 0 {
		return false

	}
	if phone[0] == '+' {
		return true
	}
	return false
}

func transformKazNumber(clearedPhone string) string {
	if !IsKazakhstanPhone(clearedPhone) {
		return clearedPhone
	}

	if len(clearedPhone) == 11 {
		if clearedPhone[0] == '8' {
			clearedPhone = "7" + clearedPhone[1:]
		}
		if code, ok := oldToNewCode[clearedPhone[1:4]]; ok {
			clearedPhone = clearedPhone[:1] + code + clearedPhone[4:]
		}
	}

	if len(clearedPhone) == 10 {
		if code, ok := oldToNewCode[clearedPhone[0:3]]; ok {
			clearedPhone = clearedPhone[:0] + code + clearedPhone[3:]
		}
		clearedPhone = "7" + clearedPhone
	}

	if len(clearedPhone) != 11 {
		return clearedPhone
	}

	return clearedPhone
}

// trimMobilePhone удаляет из строки символы +()-–-+b и пробелы
func trimMobilePhone(input string) string {
	// Итерируем по строке и удаляем символы
	for _, char := range charsToRemove {
		input = strings.ReplaceAll(input, string(char), "")
	}

	return input
}

func IsKazakhstanPhone(phone string) bool {
	if len(phone) == 11 && (string(phone[0]) == "8" || string(phone[0]) == "7") {
		code := phone[1:4]
		exist := false
		for _, s := range KazakhstanOperatorCode {
			if s == code {
				exist = true
			}
		}
		if exist {
			return true
		}
	}

	if len(phone) == 10 {
		code := phone[0:3]
		exist := false
		for _, kzCode := range KazakhstanOperatorCode {
			if kzCode == code {
				exist = true
			}
		}
		if exist {
			return true
		}
	}

	return false
}

func IsRussianPhone(phone string) bool {
	if len(phone) == 11 && (string(phone[0]) == "8" || string(phone[0]) == "7") {
		code := phone[1:3]
		if code == "89" || code == "79" {
			return true
		}
	}
	if len(phone) == 10 {
		code := []rune(phone)[0]
		if code == '9' {
			return true
		}
	}

	return false
}

func transformRussianNumber(phone string) string {
	if !IsRussianPhone(phone) {
		return phone
	}

	if len(phone) == 11 && string(phone[0]) == "8" {
		rPhon := []rune(phone)
		rPhon[0] = '7'
		return string(rPhon)
	}
	if len(phone) == 10 && string(phone[0]) == "9" {
		return "7" + phone
	}
	return phone
}

func ValidateBasicPhone(phone string) (string, bool) {
	if !IsValidPhone(phone) {
		return phone, false
	}

	firstPlus := isFirstPlus(phone)

	trimmedPhone := trimMobilePhone(phone)

	if firstPlus {
		trimmedPhone = "+" + trimmedPhone
	}
	return trimmedPhone, true
}
