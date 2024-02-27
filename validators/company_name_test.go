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
	}

	for _, tCase := range tests {
		valid := IsValidCompanyName(tCase.input)
		assert.Equal(t, tCase.valid, valid, "input: %s", tCase.input)

	}

}
