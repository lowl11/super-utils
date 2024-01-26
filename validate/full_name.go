package validate

import "regexp"

var (
	_validNameFormat             = regexp.MustCompile(`^[A-Za-zА-Яа-яӘәІіҢңҒғҮүҰұҚқӨөҺһЁё\s\-\.]+$`)
	_validCyrillicFullNameFormat = regexp.MustCompile(`[А-Яа-я]`)
	_validLatinFullNameFormat    = regexp.MustCompile(`[A-Za-z]`)
)

func CustomerName(name string) bool {
	// Проверяем, что поле Имя или Фамилия заполнены, иначе это ошибка
	if len(name) == 0 {
		return false
	}

	// Проверяем формат Имени, Фамилии и Отчества (допустимы буквы латиницы и кириллицы, пробелы, дефисы, точки)
	if len(name) != 0 && !_validNameFormat.MatchString(name) {
		return false
	}

	// Проверяем наличие букв кириллицы и латиницы в Имени, Фамилии и Отчестве
	if len(name) != 0 {
		if containsCyrillic := _validCyrillicFullNameFormat.MatchString(name); containsCyrillic {
			containsLatin := _validLatinFullNameFormat.MatchString(name)
			if containsCyrillic && containsLatin {
				return false
			}
		}
	}

	return true
}
