package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputs string) (string, error) {
	var (
		rep      int
		prevchar string
		newstr   string
		out      strings.Builder
		errout   error
	)
	for i, r := range inputs {
		rep = 1
		if unicode.IsDigit(r) && i == 0 {
			errout = ErrInvalidString
			break
		}
		if unicode.IsDigit(r) && i > 0 {
			if _, err := strconv.Atoi(prevchar); err == nil {
				errout = ErrInvalidString
				break
			}
			rep, _ = strconv.Atoi(string(r))
			if rep > 1 {
				rep--
			} else if rep == 0 {
				newstr = strings.TrimSuffix(out.String(), prevchar)
				out.Reset()
				out.WriteString(newstr)
			}
		} else {
			prevchar = string(r)
		}
		fmt.Fprintf(&out, "%s", strings.Repeat(prevchar, rep))
		prevchar = string(r)
	}
	return out.String(), errout
}
