package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoConvert(input string) (string, error) {
	if isMorseCode(input) {
		return morse.ToText(input), nil
	}
	return morse.ToMorse(input), nil
}

func isMorseCode(input string) bool {
	for _, char := range input {
		if char != '-' && char != '.' && char != ' ' && char != '\n' && char != '\r' {
			return false
		}
	}
	return true
}
