package validate

import "strconv"

var (
	// Веса разрядов
	_iinWeights  = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	_iinWeights2 = []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2}
)

func IIN(value string) bool {
	// Проверяем, что длина ИИН равна 12 символам
	if len(value) != 12 {
		return false
	}

	// Проверяем, что ИИН содержит только цифры
	if _, err := strconv.Atoi(value); err != nil {
		return false
	}

	// Проверяем формат первых 6 символов
	year, _ := strconv.Atoi(value[0:2])
	month, _ := strconv.Atoi(value[2:4])
	day, _ := strconv.Atoi(value[4:6])

	if year < 0 || year > 99 || month < 1 || month > 12 || day < 1 || day > 31 {
		return false
	}

	// Инициализируем сумму
	sum := 0
	sum2 := 0

	// Проходим по каждому разряду ИИН и вычисляем сумму
	for i := 0; i < 11; i++ {
		digit, _ := strconv.Atoi(string(value[i]))
		sum += digit * _iinWeights[i]
		sum2 += digit * _iinWeights2[i]
	}

	a12Expected, _ := strconv.Atoi(string(value[11]))
	return sum2%11 == a12Expected
}
