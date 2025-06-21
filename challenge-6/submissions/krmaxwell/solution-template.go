// Package challenge6 contains the solution for Challenge 6.
package challenge6

import "strings"

// Add any necessary imports here

// CountWordFrequency takes a string containing multiple words and returns
// a map where each key is a word and the value is the number of times that
// word appears in the string. The comparison is case-insensitive.
//
// Words are defined as sequences of letters and digits.
// All words are converted to lowercase before counting.
// All punctuation, spaces, and other non-alphanumeric characters are ignored.
//
// For example:
// Input: "The quick brown fox jumps over the lazy dog."
// Output: map[string]int{"the": 2, "quick": 1, "brown": 1, "fox": 1, "jumps": 1, "over": 1, "lazy": 1, "dog": 1}
func CountWordFrequency(text string) map[string]int {
	text = strings.ToLower(text)
	var newText strings.Builder
	for _, r := range text {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || (r == ' ') {
			newText.WriteRune(r)
		} else if r == '\'' {
			continue
		} else {
			newText.WriteRune(' ')
		}
	}
	words := strings.Fields(newText.String())

	wordFrequency := make(map[string]int)
	for i := range words {
		count, exists := wordFrequency[words[i]]
		if exists {
			wordFrequency[words[i]] = count + 1
		} else {
			wordFrequency[words[i]] = 1
		}
	}
	return wordFrequency
}
