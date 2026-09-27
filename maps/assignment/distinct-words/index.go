// Complete the countDistinctWords function using a map. It should take a slice of strings and return the total count of distinct words across all the strings. Assume words are separated by spaces. Casing should not matter. (e.g., "Hello" and "hello" should be considered the same word).
package main

import (
	"strings"
)

func countDistinctWords(messages []string) int {
	// ?
	wordMap := make(map[string]int)
	for _, message := range messages {
		for word := range strings.FieldsSeq(strings.ToLower(message)) {
			wordMap[word]++
		}
	}
	return len(wordMap)
}
