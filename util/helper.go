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

func IntToBool(i string) bool {
	if i == "1" {
		return true
	} else if i == "2" {
		return false
	}
	// Return a default value or handle other cases if necessary
	return false
}

func IntToStringBool(i string) string {
	if i == "1" {
		return "true"
	} else if i == "2" {
		return "false"
	}
	// Return a default value or handle other cases if necessary
	return "false"
}

func RemoveExtraSpaces(input string) string {
	// Разбиваем строку на слова, удаляем пустые элементы
	words := strings.Fields(input)
	// Собираем слова обратно с одним пробелом между ними
	result := strings.Join(words, " ")
	return result
}
