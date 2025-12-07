// Package iteration repeats a character multiple times.
package iteration

import "strings"

// Strings in Go are immutable. So we can't modify them in place.
//func Repeat(character string) string {
//	var repeated string
//	for i := 0; i < repeatCount; i++ {
//		repeated += character
//	}
//	return repeated
//}

func Repeat(character string, times int) string {
	var repeated strings.Builder
	for i := 0; i < times; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
