package leftpad

import (
	"errors"
	"strings"
)

// PadLeft takes the input string src and returns a string of length 'length'
// padded by the character padding
func PadLeft(src, padding string, length int) (string, error) {

	if length < len(src) {
		return "", errors.New("Target length is less than the length of the" +
			"origional string.")
	}

	if len(padding) > 1 {
		return "", errors.New("Padding must be a single character.")
	}

	return strings.Repeat(padding, length-len(src)) + src, nil

}
