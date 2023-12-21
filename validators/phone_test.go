package validators

import (
	"testing"
)

func TestValidateMobilePhone(t *testing.T) {
	type expected struct {
		phone string
		valid bool
	}
	//
	testCases := []struct {
		input    string
		expected expected
	}{

		{"7714735008", expected{
			phone: "+77714735008",
			valid: true,
		}},
		{"87714735008", expected{
			phone: "+77714735008",
			valid: true,
		}},
		{"883256582351516546", expected{
			phone: "883256582351516546",
			valid: true,
		}},
		{"7014735008", expected{
			phone: "+77014735008",
			valid: true,
		}},
		{"+81-80-46500000", expected{
			phone: "+818046500000",
			valid: true,
		}}, // Верная строка: пример японского номера
		{"8-705-123-4567", expected{
			phone: "+77051234567",
			valid: true,
		}}, // валидный кз номер
		{"+7-305-1111111", expected{
			phone: "+73051111111",
			valid: true,
		}}, // Верная строка: старый казахстанскмй номер
		{"(123) 456-7890", expected{
			phone: "+1234567890",
			valid: true,
		}}, // Верная строка: скобки, пробелы и тире
		{"+1-800-555-1234", expected{
			phone: "+18005551234",
			valid: true,
		}}, // Верная строка: знак плюса и тире
		{"883256582351516546", expected{
			phone: "883256582351516546",
			valid: true,
		}}, // Верная строка: иностранный номер начинающийся на 8
		{"+7-105-1142306", expected{
			phone: "+7-105-1142306",
			valid: false,
		}}, // Верная строка: несуществующий казахстанскмй номер
		{"+1+800-555-1234", expected{
			phone: "+1+800-555-1234",
			valid: false,
		}}, // Недопустимая строка: буквы и дефис
		{"+1+5352-555-1234", expected{
			phone: "+1+5352-555-1234",
			valid: false,
		}}, // Недопустимая строка: много с
		{"1+800-555-1234", expected{
			phone: "1+800-555-1234",
			valid: false,
		}}, // Недопустимая строка: плюч не в начале
		{"123 + (45 - 67)", expected{
			phone: "123 + (45 - 67)",
			valid: false,
		}}, // Недопустимая строка: + не в начале
		{"(555) 123-4567 ext. 890", expected{
			phone: "(555) 123-4567 ext. 890",
			valid: false,
		}}, // Недопустимая строка: буквы
		{"123 456", expected{
			phone: "123 456",
			valid: false,
		}}, // Верная строка: пробелы
		{"ABC-DEF-GHI", expected{
			phone: "ABC-DEF-GHI",
			valid: false,
		}}, // Недопустимая строка: буквы и дефис

	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			phone, valid := ValidateMobilePhone(testCase.input)
			if phone != testCase.expected.phone && valid == testCase.expected.valid {
				t.Errorf("Для строки '%s' ожидается phone %v valid %v, но получено phone %s, valid %v",
					testCase.input, testCase.expected.phone, testCase.expected.valid, phone, valid)
			}
		})
	}

}

func TestValidatePhone(t *testing.T) {
	type expected struct {
		phone string
		valid bool
	}
	testCases := []struct {
		input    string
		expected expected
	}{
		{"+81-80-46500000", expected{
			phone: "+818046500000",
			valid: true,
		}}, // Верная строка: пример японского номера
		{"+7-305-1111111", expected{
			phone: "+73051111111",
			valid: true,
		}}, // Верная строка: старый казахстанскмй номер
		{"(123) 456-7890", expected{
			phone: "1234567890",
			valid: true,
		}}, // Верная строка: скобки, пробелы и тире
		{"+1-800-555-1234", expected{
			phone: "+18005551234",
			valid: true,
		}}, // Верная строка: знак плюса и тире
		{"+7-105-1142306", expected{
			phone: "+7-105-1142306",
			valid: false,
		}}, // Верная строка: несуществующий казахстанскмй номер
		{"+1+800-555-1234", expected{
			phone: "+1+800-555-1234",
			valid: false,
		}}, // Недопустимая строка: буквы и дефис
		{"+1+5352-555-1234", expected{
			phone: "+1+5352-555-1234",
			valid: false,
		}}, // Недопустимая строка: много с
		{"1+800-555-1234", expected{
			phone: "1+800-555-1234",
			valid: false,
		}}, // Недопустимая строка: плюч не в начале
		{"123 + (45 - 67)", expected{
			phone: "123 + (45 - 67)",
			valid: false,
		}}, // Недопустимая строка: + не в начале
		{"(555) 123-4567 ext. 890", expected{
			phone: "(555) 123-4567 ext. 890",
			valid: false,
		}}, // Недопустимая строка: буквы
		{"123 456", expected{
			phone: "123 456",
			valid: false,
		}}, // Верная строка: пробелы
		{"ABC-DEF-GHI", expected{
			phone: "ABC-DEF-GHI",
			valid: false,
		}}, // Недопустимая строка: буквы и дефис
		{"322-22-22-333", expected{
			phone: "3222222333",
			valid: true,
		}}, // Недопустимая строка: буквы и дефис

	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			phone, valid := ValidateBasicPhone(testCase.input)
			if phone != testCase.expected.phone && valid == testCase.expected.valid {
				t.Errorf("Для строки '%s' ожидается phone %v valid %v, но получено phone %s, valid %v",
					testCase.input, testCase.expected.phone, testCase.expected.valid, phone, valid)
			}
		})
	}

}
