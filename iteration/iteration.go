package iteration

import "strings"

func Repeat(character string, repetitions int) string {
	var repeated strings.Builder
	for i := 0; i < repetitions; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
