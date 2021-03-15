package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	check := make([]rune, 0, len(s))
	for _, r := range s {
		check = append(check, r)
		isDigit := unicode.IsDigit(r)
		if isDigit {
			if len(check) == 1 {
				return "", ErrInvalidString
			}
			ll := check[len(check)-2]
			CheckDigit := unicode.IsDigit(ll)
			if CheckDigit {
				return "", ErrInvalidString
			}
		}
	}
	l := make([]rune, 0, len(s))
	var letter string
	for _, r := range s {
		isDigit := unicode.IsDigit(r)
		if !isDigit {
			l = append(l, r)
		}
		if isDigit && r != '0' {
			if string(l) == "" {
				return "", ErrInvalidString
			}
			q, _ := strconv.Atoi(string(r))
			ll := l[len(l)-1]
			letter += strings.Repeat(string(ll), q-1)
		}
		if !isDigit {
			letter += string(r)
		}
		if r == '0' {
			var l []rune
			for _, r := range letter {
				l = append(l, r)
			}
			l = l[:len(l)-1]
			letter = string(l)
		}
	}
	return letter, nil
}
