// Complete the removeProfanity function.

// It should use the strings.ReplaceAll function to replace all instances of the following words in the input message with asterisks.

// Word	Replacement
// fubb	****
// shiz	****
// witch	*****
// It should update the value in the pointer and return nothing.

package main

import (
	"strings"
)

func removeProfanity(message *string) {
	// ?
	*message = strings.ReplaceAll(*message, "fubb", "****")
	*message = strings.ReplaceAll(*message, "shiz", "****")
	*message = strings.ReplaceAll(*message, "witch", "*****")
}
