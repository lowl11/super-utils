package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateNumberOfRegistration(t *testing.T) {
	type TestCase struct {
		input string
		value string
		valid bool
	}
	tests := []TestCase{
		{"СЕРИЯ 10915 № 0145240", "СЕРИЯ 10915 № 0145240", true},
		{"asd", "asd", true},
		{"asd-12", "asd-12", true},
		{"ыфвф", "ыфвф", true},
		{"ыфвф-12", "ыфвф-12", true},
		{"こんにちは 你好", "こんにちは 你好", false},
	}
	{
		for _, tCase := range tests {
			res, valid := NumberOfRegistration(tCase.input)
			assert.Equal(t, tCase.value, res)
			assert.Equal(t, tCase.valid, valid)

		}
	}

}
