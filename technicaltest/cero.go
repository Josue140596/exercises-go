package technicaltest

import "strings"

// In this test we need to know if a string is a palindrome
// A palindrome is a word that reads the same backward as forward.
// For instance: "Anita lava la tina", "seres", "Rapar"

// Another important thing to mention, this implementation is case-insensitive
// and ignores spaces.

// Case-insensitive means that "Hello," "hello," and "HELLO" would be considered the same.
func TTPalindrome() {
	println("Anita lava la tina", isPalindrome("Anita lava la tina"))
	println("Rapar", isPalindrome("Rapar"))
	println("PaLiNDroMe", isPalindrome("PaLiNDroMe"))

}

func isPalindrome(word string) bool {
	//First we need to clean our string deleting spaces
	//ReplaceAll find a space an replace by ""
	cleanWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))
	//anitalavalatina
	left := 0
	right := len(cleanWord) - 1
	for left < right {
		if cleanWord[left] != cleanWord[right] {
			return false
		}
		left++
		right--
	}
	return true
}
