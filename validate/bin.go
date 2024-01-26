package validate

import (
	"strconv"
)

func BIN(bin string) bool {
	// Проверка длины
	if len(bin) != 12 {
		return false
	}
	// Проверка первых 4 символов
	year, err := strconv.Atoi(bin[0:2])
	if err != nil || year < 0 || year > 99 {
		return false
	}
	month, err := strconv.Atoi(bin[2:4])
	if err != nil || month < 1 || month > 12 {
		return false
	}
	// Проверка пятого символа
	fifthDigit, err := strconv.Atoi(string(bin[4]))
	if err != nil || fifthDigit < 4 || fifthDigit > 9 {
		return false
	}
	// Проверка шестого символа
	sixthDigit, err := strconv.Atoi(string(bin[5]))
	if err != nil || (sixthDigit <= 0 && sixthDigit >= 4) {
		return false
	}

	return true
}
