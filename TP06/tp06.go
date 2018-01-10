package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(areAnagrams("this is not", "an anagram"))
	fmt.Println(areAnagrams("mary", "army"))
	fmt.Println(areAnagrams("silent", "listen"))
	fmt.Println(areAnagrams("Madam Curie", "Radium came"))
	fmt.Println(areAnagrams("william Shakespeare", "I am a weakish speller"))
	fmt.Println(areAnagrams("😁😞", "😞😁"))
}

func areAnagrams(a, b string) bool {
	a = normalize(a)
	b = normalize(b)

	if len(a) != len(b) {
		return false
	}

	for _, c := range a {
		var i int
		if i = strings.IndexRune(b, c); i == -1 {
			return false
		}
		b = removeRuneAtIndex(b, i, utf8.RuneLen(c))
	}
	return true
}

func normalize(s string) string {
	return removeWhiteSpaces(toLower(s))
}

func toLower(s string) string {
	// Add the code to return the lower case version of s
	// Hint: look in the standard library strings package
	return strings.ToLower(s)
}

func removeWhiteSpaces(s string) string {
	// Add the code to return s without any white space
	// Hint: use functions Fields and Join of the strings package
        return strings.Replace(s, " ", "", -1)
}

func removeRuneAtIndex(s string, i int, l int) string {
	// Add the code to return a new string that is
	// identical to s except that the rune at index i
	// with length l has been removed
	return s[:i]+s[i+l:]
}
