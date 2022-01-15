package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const (
	Backslash = 92
	Step      = 1
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var res strings.Builder
	ra := []rune(str)
	for i := 0; i < len(ra); i++ {
		if ra[i] == Backslash {
			i++
			res.WriteRune(ra[i])
			continue
		}

		if n, err := strconv.Atoi(string(ra[i])); err == nil {
			if i == 0 || (i > 0 && unicode.IsDigit(ra[i-1]) && (i > 1 && ra[i-2] != Backslash)) {
				return "", ErrInvalidString
			}

			if n == 0 {
				var strRes = res.String()
				strRes = strRes[:len(strRes)-1]
				res.Reset()
				res.WriteString(strRes)
				continue
			}

			res.WriteString(strings.Repeat(string(ra[i-1]), n-Step))
			continue
		}

		res.WriteRune(ra[i])
	}
	return res.String(), nil
}
