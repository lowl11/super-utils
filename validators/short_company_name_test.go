package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformFullNameToShortName(t *testing.T) {
	type TestCase struct {
		expected string
		input    string
		valid    bool
	}

	// "АО", FullName: "Акционерное общество"
	tests := []TestCase{
		{"ТооКазахстан-2030", "ТооКазахстан-2030", true},
		{"some", "some", true},
		{"ТОО Казахстан-2030", "Товарищество с ограниченной ответственностью Казахстан-2030", true},
		{"АО Казахстан-2030", "Акционерное общество Казахстан-2030", true},
		{"АОЗТ Казахстан-2030", "Акционерное общество закрытого типа Казахстан-2030", true},
		{"ГКП Казахстан-2030", "Государственное коммунальное предприятие Казахстан-2030", true},
		{"ГКП на ПХВ Казахстан-2030", "Государственное коммунальное предприятие на праве хозяйственного ведения Казахстан-2030", true},
		{"LP Terminal", "Limited Partnership Terminal", true},
		{"LP", "LP ", true},
		{"ТОО$&+*", "ТОО$&+*", false},
	}
	for _, tCase := range tests {
		result, isValid := FullNameToAbbreviation(tCase.input)
		assert.Equal(t, tCase.expected, result, "input: %s", tCase.input)
		assert.Equal(t, tCase.valid, isValid, "input: %s", tCase.input)

	}

}
