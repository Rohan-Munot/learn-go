// Assignment
// Implement the formatter interface with a method format that returns a formatted string.
// Define structs that satisfy the formatter interface: plainText, bold, code.
// The structs must all have a message field of type string
// plainText should return the message as is.
// bold should wrap the message in two asterisks (**) to simulate bold text (e.g., **message**).
// code should wrap the message in a single backtick (`) to simulate inline code (e.g., `message`)

package main

import "fmt"

type formatter interface {
	format() string
}

type plainText struct {
	message string
}

type bold struct {
	message string
}

type code struct {
	message string
}

func (p plainText) format() string {
	return p.message
}

func (b bold) format() string {
	return fmt.Sprintf("**%s**", b.message)
}

func (c code) format() string {
	return fmt.Sprintf("`%s`", c.message)
}

// Don't Touch below this line

func sendMessage(format formatter) string {
	return format.format() // Adjusted to call Format without an argument
}
