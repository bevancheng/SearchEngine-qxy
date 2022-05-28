package utils

import "regexp"

func RemovePunctuation(text string) string {
	reg := regexp.MustCompile(`\p{P}+`)
	return reg.ReplaceAllString(text, "")
}
func RemoveSpace(text string) string {
	reg := regexp.MustCompile(`\s+`)
	return reg.ReplaceAllString(text, "")
}
