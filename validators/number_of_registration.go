package validators

import "regexp"

func ValidateNumberOfRegistration(str string) (string, bool) {
	// Допустимые символы только: цифры, буквы латиницы и кириллицы,
	//пробелы, дефисы/тире, точки, скобки, слэш, двоеточие. Содержит либо только буквы кириллицы
	//(без букв латиницы), либо должно содержать только буквы латиницы без букв кириллицы).
	re := regexp.MustCompile(`^([А-Яа-яЁё№\s\d-.()/:]+|[A-Za-z№\s\d-.()/:]+)$`)

	// Проверка строки на соответствие регулярному выражению
	if re.MatchString(str) {
		return str, true
	}
	return str, false
}
