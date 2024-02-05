package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateFirstName(t *testing.T) {
	type TestCase struct {
		input string
		value string
		valid bool
	}
	tests := []TestCase{
		{"Саня", "Саня", true},
		{"саня", "саня", true},
		{"   саня  ", "саня", true},
		{"Жан-Клод", "Жан-Клод", true},
		{"Жан Клод", "Жан Клод", true},
		{"Черепашка—Ниндзя", "Черепашка—Ниндзя", true},
		{"П.Томас", "П.Томас", true},
	}
	{
		for _, tCase := range tests {
			res, valid := ValidateFirstName(tCase.input)
			assert.Equal(t, tCase.value, res, "input: %s", tCase.input)
			assert.Equal(t, tCase.valid, valid, "input: %s", tCase.input)

		}
	}

}
