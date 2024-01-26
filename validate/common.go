package validate

import "fmt"

func Required(value any) bool {
	if value == nil {
		return false
	}

	if len(fmt.Sprintf("%v", value)) == 0 {
		return false
	}

	return true
}
