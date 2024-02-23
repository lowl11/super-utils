package validators

import "testing"

func TestIsValidIinNegative(t *testing.T) {
	type expected struct {
		Valid bool
	}
	//
	testCases := []struct {
		input    string
		expected expected
	}{
		//значение ИИН 13 символов
		{"8011284008299", expected{
			Valid: false,
		}},
		//значение ИИН 11 символов
		{"80112840082", expected{
			Valid: false,
		}},
		//значение ИИН несуществующий месяц (3-ий и 4-ый символы)
		{"803328400829", expected{
			Valid: false,
		}},
		//значение ИИН несуществующая дата (5-ый и 6-ой символы)
		{"801133400829", expected{
			Valid: false,
		}},
		//значение ИИН настоящие, кроме последней цифры
		{"821128400836", expected{
			Valid: false,
		}},
		//значение ИИН буквами или символами
		{"йййцццуууккк", expected{
			Valid: false,
		}},
		//значение не заполнено
		{"", expected{
			Valid: false,
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			valid, err := IsValidIin(testCase.input)
			if valid != testCase.expected.Valid &&
				err == nil {
				t.Errorf("Для строки '%s' valid %v, но получено valid %v",
					testCase.input, testCase.expected.Valid, valid)
			}
		})
	}

}
