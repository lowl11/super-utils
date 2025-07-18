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
		// Валидные: Примеры валидации номеров согласно правилам идентификации 010, 020 для мобильных номеров
		{"7714735008", expected{phone: "+77714735008", valid: true}},
		{"87714735008", expected{phone: "+77714735008", valid: true}},
		{"883256582351516546", expected{phone: "+883256582351516546", valid: true}},
		{"7014735008", expected{phone: "+77014735008", valid: true}},
		{"7004735008", expected{phone: "+77004735008", valid: true}},
		{"87004735008", expected{phone: "+77004735008", valid: true}},
		{"+77004735008", expected{phone: "+77004735008", valid: true}},
		{"+81-80-46500000", expected{phone: "+818046500000", valid: true}},
		{"8-705-123-4567", expected{phone: "+77051234567", valid: true}},
		{"+7-305-1111111", expected{phone: "+77051111111", valid: true}},
		{"(123) 456-7890", expected{phone: "+1234567890", valid: true}},
		{"+1-800-555-1234", expected{phone: "+18005551234", valid: true}},

		// Ошибка: недопустимые символы (буквы, спецсимволы)
		{"ABC-DEF-GHI", expected{phone: "ABC-DEF-GHI", valid: false}},                         // буквы и дефис
		{"(555) 123-4567 ext. 890", expected{phone: "(555) 123-4567 ext. 890", valid: false}}, // буквы и спец. добавка
		{"№93759!=;", expected{phone: "№93759!=;", valid: false}},                             // спецсимволы

		// Ошибка: неправильное расположение символа '+'
		{"+1+800-555-1234", expected{phone: "+1+800-555-1234", valid: false}},   // более одного '+'
		{"+1+5352-555-1234", expected{phone: "+1+5352-555-1234", valid: false}}, // более одного '+'
		{"1+800-555-1234", expected{phone: "1+800-555-1234", valid: false}},     // '+' не в начале
		{"123 + (45 - 67)", expected{phone: "123 + (45 - 67)", valid: false}},   // '+' не в начале

		// Ошибка: некорректный формат или неполный номер
		{"123 456", expected{phone: "123 456", valid: false}},        // слишком короткий номер / неформат
		{" ", expected{phone: " ", valid: false}},                    // только пробел
		{"8701777224", expected{phone: " 8701777224", valid: false}}, // длина меньше 10 символов (и пробел спереди)

		// Ошибка: валидный формат, но несуществующий код
		{"+7-105-1142306", expected{phone: "+7-105-1142306", valid: false}}, // несуществующий казахстанский код
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
		{" ", expected{
			phone: " ",
			valid: false,
		}},
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

func TestValidateLocPhone(t *testing.T) {
	type expected struct {
		phone string
		valid bool
	}
	//
	type testType struct {
		input    string
		expected expected
		comment  string
	}

	testCases := []testType{
		{
			input: "",
			expected: expected{
				phone: "",
				valid: true,
			},
			comment: "TZ1: Positive - Empty string is valid as field is not mandatory.",
		},
		{
			input: "+819012345678",
			expected: expected{
				phone: "+1234567890",
				valid: true,
			},
			comment: "check for Japan phone number",
		},
		{
			input: "+7 (700) 123-45-67",
			expected: expected{
				phone: "+77001234567",
				valid: true,
			},
			comment: "TZ3: Positive - All allowed characters (digits, plus, spaces, parens, hyphen). Valid KZ number.",
		},
		{
			input: "87051234567",
			expected: expected{
				phone: "+77051234567",
				valid: true,
			},
			comment: "TZ3: Positive - Only digits. Valid KZ 11-digit number, normalized.",
		},
		{
			input: "777 111 2233",
			expected: expected{
				phone: "+77771112233",
				valid: true,
			},
			comment: "TZ3: Positive - Digits and spaces. Valid KZ 10-digit number.",
		},
		{
			input: "(701)234-5678",
			expected: expected{
				phone: "+77012345678",
				valid: true,
			},
			comment: "TZ3: Positive - Digits, spaces, parens, hyphen. Valid KZ 10-digit number.",
		},
		{
			input: "+7 (700) 123-45-67!",
			expected: expected{
				phone: "+7 (700) 123-45-67!",
				valid: false,
			},
			comment: "TZ3: Negative - Contains disallowed character '!'",
		},
		{
			input: "8,705,123,45,67",
			expected: expected{
				phone: "8,705,123,45,67",
				valid: false,
			},
			comment: "TZ3: Negative - Contains disallowed character ','",
		},
		{
			input: "abc1234567",
			expected: expected{
				phone: "abc1234567",
				valid: false,
			},
			comment: "TZ3: Negative - Contains letters 'abc'",
		},
		{
			input: "700.123.4567",
			expected: expected{
				phone: "700.123.4567",
				valid: false,
			},
			comment: "TZ3: Negative - Contains disallowed character '.'",
		},
		{
			input: "+7 (700) 123–45–67",
			expected: expected{
				phone: "+7 (700) 123–45–67",
				valid: false,
			},
			comment: "TZ3: Negative - Contains en dash '–' (U+2013) instead of hyphen-minus '-' (U+002D) (assuming only hyphen-minus is allowed by `allowedCharsPattern`)",
		},
		{
			input: "+77001234567",
			expected: expected{
				phone: "+77001234567",
				valid: true,
			},
			comment: "TZ4: Positive - Starts with '+' and is a valid KZ number.",
		},
		{
			input: "87001234567",
			expected: expected{
				phone: "+77001234567",
				valid: true,
			},
			comment: "TZ4: Positive - Starts with digit '8' and is a valid KZ number (normalized).",
		},
		{
			input: " (700) 1234567",
			expected: expected{
				phone: "+77001234567",
				valid: true,
			},
			comment: "TZ4: Positive - Starts with space, but first significant char is '(' then digit '7'. Valid KZ number.",
		},
		{
			input: "-(700)1234567",
			expected: expected{
				phone: "-(700)1234567",
				valid: false,
			},
			comment: "TZ4: Negative - Starts with '-', not plus or digit.",
		},
		{
			input: " 7+7001234567",
			expected: expected{
				phone: " 7+7001234567",
				valid: false,
			},
			comment: "TZ4: Negative - Contains '+', but not as the first character.",
		},
		{
			input: "700+1234567",
			expected: expected{
				phone: "700+1234567",
				valid: false,
			},
			comment: "TZ4: Negative - '+' in the middle of the string.",
		},
		{
			input: "7001234567+",
			expected: expected{
				phone: "7001234567+",
				valid: false,
			},
			comment: "TZ4: Negative - '+' at the end of the string.",
		},
		{
			input: "++77001234567",
			expected: expected{
				phone: "++77001234567",
				valid: false,
			},
			comment: "TZ4: Negative - Multiple '+' characters.",
		},
		{
			input: " (abc) 1234567",
			expected: expected{
				phone: " (abc) 1234567",
				valid: false,
			},
			comment: "TZ3 & TZ4: Negative - Starts with space, then '(', then non-digit 'a'. Fails allowed chars (TZ3) and first significant char rule (TZ4).",
		},
		{
			input: "+ () -",
			expected: expected{
				phone: "+ () -",
				valid: false,
			},
			comment: "TZ4: Negative - Only non-digit allowed characters, results in no extracted digits.",
		},
		{
			input: "() - ",
			expected: expected{
				phone: "() - ",
				valid: false,
			},
			comment: "TZ4: Negative - Starts with '(', not plus or digit, and no digits extracted.",
		},
		{
			input: "7051234567",
			expected: expected{
				phone: "+77051234567",
				valid: true,
			},
			comment: "TZ5.1: Positive - 10 digits, valid KZ operator code '705'.",
		},
		{
			input: "+7 (705) 123-45-67",
			expected: expected{
				phone: "+77051234567",
				valid: true,
			},
			comment: "TZ5.1: Positive - Formatted, 10 digits effectively '705...', valid KZ operator code.",
		},
		{
			input: "3051234567",
			expected: expected{
				phone: "+73051234567",
				valid: true,
			},
			comment: "TZ5.1: Positive - 10 digits, valid old KZ operator code '305'.",
		},
		{
			input: "+7-305-123-45-67",
			expected: expected{
				phone: "+7-305-123-45-67",
				valid: false,
			},
			comment: "TZ5.1 example clarification / TZ5.2: Input +7-305..., digits 7305... (11 digits). Starts '7', not '77', not '79' (it's '73'). Fails TZ5.2 (continuation) rule.",
		},
		{
			input: "9001234567",
			expected: expected{
				phone: "+9001234567",
				valid: true,
			},
			comment: "TZ5.1 Negative (not KZ code), falls to TZ5.4: 10 digits, but '900' is not a listed KZ code. Valid as 'other'.",
		},
		{
			input: "705123456",
			expected: expected{
				phone: "705123456",
				valid: false,
			},
			comment: "TZ5.1 Negative (not 10 digits). Then fails TZ5.2 (7xx not 77, not 79) due to prefix and length.",
		},
		{
			input: "77011234567",
			expected: expected{
				phone: "+77011234567",
				valid: true,
			},
			comment: "TZ5.2.1 & TZ5.2.2: Positive - KZ, 11 digits, starts '77', valid operator code '701'.",
		},
		{
			input: "87051234567",
			expected: expected{
				phone: "+77051234567",
				valid: true,
			},
			comment: "TZ5.2.1 & TZ5.2.2: Positive - KZ, 11 digits, starts '87', valid operator code '705', normalized.",
		},
		{
			input: "7701123456",
			expected: expected{
				phone: "7701123456",
				valid: false,
			},
			comment: "TZ5.2.1: Negative - KZ prefix '77', but 10 digits (not 11).",
		},
		{
			input: "870512345678",
			expected: expected{
				phone: "870512345678",
				valid: false,
			},
			comment: "TZ5.2.1: Negative - KZ prefix '87', but 12 digits (not 11).",
		},
		{
			input: "77991234567",
			expected: expected{
				phone: "77991234567",
				valid: false,
			},
			comment: "TZ5.2.2: Negative - KZ, 11 digits, '77', but invalid operator code '799'.",
		},
		{
			input: "87331234567",
			expected: expected{
				phone: "87331234567",
				valid: false,
			},
			comment: "TZ5.2.2: Negative - KZ, 11 digits, '87', but invalid operator code '733'.",
		},
		{
			input: "71234567890",
			expected: expected{
				phone: "71234567890",
				valid: false,
			},
			comment: "TZ5.2 (continuation), 5.2.1 (for РФ): Negative - Starts '7', not '77', and not '79' (it's '71').",
		},
		{
			input: "79001234567",
			expected: expected{
				phone: "+79001234567",
				valid: true,
			},
			comment: "TZ5.2 (continuation), 5.2.1 & 5.2.2 (for РФ): Positive - Russian, '79', 11 digits.",
		},
		{
			input: "7900123456",
			expected: expected{
				phone: "7900123456",
				valid: false,
			},
			comment: "TZ5.2 (continuation), 5.2.2 (for РФ): Negative - Russian, '79', but 10 digits (not 11).",
		},
		{
			input: "89001234567",
			expected: expected{
				phone: "+79001234567",
				valid: true,
			},
			comment: "TZ5.3: Positive - Russian '89', 11 digits, normalized to '+79...'",
		},
		{
			input: "8900123456",
			expected: expected{
				phone: "8900123456",
				valid: false,
			},
			comment: "TZ5.3: Negative - Russian '89' prefix, but 10 digits (expects 11 for 89xxxxxxxxx).",
		},
		{
			input: "88001234567",
			expected: expected{
				phone: "+88001234567",
				valid: true,
			},
			comment: "TZ5.3: Positive - Starts '8', not '87', not '89' (e.g. '88'), 11 digits. No specific checks.",
		},
		{
			input: "8123456789",
			expected: expected{
				phone: "+8123456789",
				valid: true,
			},
			comment: "TZ5.3: Positive - Starts '8', not '87', not '89', 10 digits. No specific checks.",
		},
		{
			input: "1234567890",
			expected: expected{
				phone: "+1234567890",
				valid: true,
			},
			comment: "TZ5.4: Positive - Starts not '7' or '8' (e.g. '1'). No specific checks.",
		},
		{
			input: "305123456",
			expected: expected{
				phone: "+305123456",
				valid: true,
			},
			comment: "TZ5.4: Positive - Starts '3' (old KZ prefix), 9 digits. Not 10 for TZ5.1, falls to TZ5.4. No specific checks.",
		},
		{
			input: "8-305-123-4567",
			expected: expected{
				phone: "+83051234567",
				valid: true,
			},
			comment: "TZ5.1 example interpretation / TZ5.3: Input 8-305..., digits 8305... (11 digits). Starts '8', not '87'. Valid under TZ5.3.",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			phone, valid := ValidateMobilePhone(testCase.input)
			if phone != testCase.expected.phone && valid == testCase.expected.valid {
				t.Errorf("Для строки '%s' ожидается phone %v valid %v, но получено phone %s, valid %v\n Comment: %s",
					testCase.input, testCase.expected.phone, testCase.expected.valid, phone, valid, testCase.comment)
			}
		})
	}

}
