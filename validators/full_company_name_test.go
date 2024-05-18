package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCompanyName(t *testing.T) {
	type TestCase struct {
		input string
		valid bool
	}
	tests := []TestCase{
		{"Казахстан-2030", true},
		{"Жан-Клод", true},
		{"ЧерепашқҢ—Ниндзя", true},
		{"әғқңөұүһі", true},
		{"ЧерепашқҢ/Ниндзя", true},
		{"ЧерепашқҢ\\Ниндзя", true},
		{"ТОО$&+*", false},
		{"Name «2024» / , . название ", true},
		{"RBS BANK ROMANIA S.A.", true},
	}

	for _, tCase := range tests {
		valid := IsValidCompanyName(tCase.input)
		assert.Equal(t, tCase.valid, valid, "input: %s", tCase.input)

	}

}

func TestTransformateCompanyName(t *testing.T) {
	type TestCase struct {
		input    string
		expected string
		valid    bool
	}

	// "АО", FullName: "Акционерное общество"
	tests := []TestCase{
		{"ТооКазахстан-2030", "ТооКазахстан-2030", true},
		{"some", "some", true},
		{"ТОО Казахстан-2030", "Товарищество с ограниченной ответственностью Казахстан-2030", true},
		{"тоо Казахстан-2030", "Товарищество с ограниченной ответственностью Казахстан-2030", true},
		{"АО Казахстан-2030", "Акционерное общество Казахстан-2030", true},
		{"АОЗТ Казахстан-2030", "Акционерное общество закрытого типа Казахстан-2030", true},
		{"ГКП Казахстан-2030", "Государственное коммунальное предприятие Казахстан-2030", true},
		{"ГКП на ПХВ Казахстан-2030", "Государственное коммунальное предприятие на праве хозяйственного ведения Казахстан-2030", true},
		{"ИП Казахстан-2030", "Индивидуальный предприниматель Казахстан-2030", true},
		{"LP Terminal", "Limited Partnership Terminal", true},
		{"LP ", "LP", true},

		{"ТОО$&+*", "ТОО$&+*", false},
	}
	for _, tCase := range tests {
		result, isValid := AbbreviationToFullName(tCase.input)
		assert.Equal(t, tCase.expected, result, "input: %s", tCase.input)
		assert.Equal(t, tCase.valid, isValid, "input: %s", tCase.input)

	}

}
