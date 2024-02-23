package validators

import (
	"fmt"
	"strconv"
)

// IsValidIin проверяет корректность ИИН
// error - причина некорректности ИИН
func IsValidIin(iin string) (bool, error) {
	// Проверяем, что длина ИИН равна 12 символам
	if len(iin) != 12 {
		return false, fmt.Errorf("некорректная длина ИИН. ИИН должен содержать 12 символов")
	}

	// Проверяем, что ИИН содержит только цифры
	if _, err := strconv.Atoi(iin); err != nil {
		return false, fmt.Errorf("ИИН должен содержать только цифры")
	}

	// Веса разрядов
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	weights2 := []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2}

	// Проверяем формат первых 6 символов
	year, _ := strconv.Atoi(iin[0:2])
	month, _ := strconv.Atoi(iin[2:4])
	day, _ := strconv.Atoi(iin[4:6])

	if year < 0 || year > 99 || month < 1 || month > 12 || day < 1 || day > 31 {
		return false, fmt.Errorf("некорректный формат первых 6 символов ИИН")
	}

	// Инициализируем сумму
	sum := 0
	sum2 := 0

	// Проходим по каждому разряду ИИН и вычисляем сумму
	for i := 0; i < 11; i++ {
		digit, _ := strconv.Atoi(string(iin[i]))
		sum += digit * weights[i]
		sum2 += digit * weights2[i]
	}

	a12Expected, _ := strconv.Atoi(string(iin[11]))
	if sum%11 == 10 {
		if sum2%11 == a12Expected {
			return true, nil
		}
	} else if sum%11 != a12Expected {
		return false, fmt.Errorf("двенадцатая цифра ИИН не соответствует расчетному значению")
	}

	return true, nil
}
