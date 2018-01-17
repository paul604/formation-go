package anagram

import (
	"fmt"
)

func ExampleAreAnagrams() {
	fmt.Println(AreAnagrams("this is not", "an anagram"))
	fmt.Println(AreAnagrams("mary", "army"))
	fmt.Println(AreAnagrams("silent", "listen"))
	fmt.Println(AreAnagrams("Madam Curie", "Radium came"))
	fmt.Println(AreAnagrams("william Shakespeare", "I am a weakish speller"))
	// Output:
	// false
	// true
	// true
	// true
	// true
}
