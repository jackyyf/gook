package isbn

import (
	"strings"
)

func Normalize(input string) (isbn string) {
	// Empty string for error, ISBN13 otherwise.
	input = strings.Replace(input, "-", "", -1)
	l := len(input)
	if l == 10 {
		if !isValid10(input) {
			return ""
		} else {
			// Normalize to 13-byte ISBN
			prefix := "978" + input[:9]
			tot := 0
			for idx, val := range prefix {
				mul := 1
				if idx&1 != 0 {
					mul = 3
				}
				tot += (int(val) - 48) * mul
			}
			prefix += string(byte(48 + (10-tot%10)%10))
			return prefix
		}
	} else if l == 13 {
		if !isValid13(input) {
			return ""
		} else {
			return input
		}
	} else {
		return ""
	}
}

func isValid10(input string) bool {
	sum := 0
	multiply := 10
	for i, v := range input {
		digit := 0
		if i == 9 && (v == 'x' || v == 'X') {
			digit = 10
		} else if v < '0' || v > '9' {
			return false
		} else {
			digit = int(v) - '0'
		}
		sum = sum + (multiply * digit)
		multiply--
	}

	return sum%11 == 0
}

func isValid13(input string) bool {
	sum := 0
	for i, v := range input {
		multiply := 1
		if i&1 != 0 {
			multiply = 3
		}
		if v < '0' || v > '9' {
			return false
		}
		digit := int(v) - '0'

		sum = sum + (multiply * digit)
	}
	return sum%10 == 0
}
