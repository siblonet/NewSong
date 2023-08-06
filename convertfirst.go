package main

import "fmt"

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func Capitalize(s string) string {
	var result string
	capitalizeNext := true

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if isLetter(ch) {
			if capitalizeNext {
				result += string(ch &^ 32) // Convert to uppercase using bit manipulation.
				capitalizeNext = false
			} else {
				result += string(ch | 32) // Convert to lowercase using bit manipulation.
			}
		} else {
			result += string(ch)
			capitalizeNext = true

		}
	}

	return result
}

func main() {
	de := Capitalize("bb$\"ZvxNq?w?i")
	fmt.Println(de)
	def := Capitalize("uNTqn@D [B,Cf")
	fmt.Println(def)
}
