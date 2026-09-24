// As part of improving security, Textio wants to enforce a new password policy. A valid password must meet the following criteria:

// At least 5 characters long but no more than 12 characters.
// Contains at least one uppercase letter.
// Contains at least one digit.
package main

import "unicode"

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}
	hasUpper := false
	hasDigit := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true
		}
		if unicode.IsUpper(char) {
			hasUpper = true
		}
	}
	return hasUpper && hasDigit
}
