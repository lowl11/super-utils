package util

import (
	"reflect"
	"strings"
)

func Contains(slice []string, value string) (bool, string) {
	for _, elem := range slice {
		if elem == value {
			return true, "1"
		}
	}
	return false, "0"
}

func ParseFieldName(f reflect.StructField) (name string, ignore bool) {
	tag := f.Tag.Get("json")

	if tag == "" {
		return f.Name, false
	}

	if tag == "-" {
		return "", true
	}

	if i := strings.Index(tag, ","); i != -1 {
		if i == 0 {
			return f.Name, false
		} else {
			return tag[:i], false
		}
	}

	return tag, false
}

func MergeBool(a, b bool) bool {
	if a == false && b == false {
		return false
	}

	return a == b
}
