package operations

import "strings"

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

func ToLowerCase(s string) string {
	return strings.ToLower(s)
}

func ReplaceSpaces(s, replacement string) string {
	return strings.ReplaceAll(s, " ", replacement)
}
