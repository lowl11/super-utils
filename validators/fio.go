package validators

import "regexp"

// ValidateFIO alfa version
func ValidateFIO(data string) (string, bool) {
	// Проверяем формат Имени, Фамилии и Отчества (допустимы буквы латиницы и кириллицы, пробелы, дефисы, точки)
	validFormat := regexp.MustCompile(`^[A-Za-zА-Яа-яҚқҢңӘәҰұҮүІіЁёҒғөӨһҺ\s\-\.]+$`)

	ok := validFormat.MatchString(data)
	if !ok {
		return data, false
	}

	return data, true
}
