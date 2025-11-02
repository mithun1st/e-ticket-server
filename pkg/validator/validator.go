package appvalidator

import (
	"regexp"
	"unicode/utf8"
)

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func IsPasswordLengthValid(p string) bool {
	len := utf8.RuneCountInString(p)
	return len >= 8
}
