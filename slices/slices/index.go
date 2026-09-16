// comapred to arrays slices are not fixed in size, slices are dynamically sized flexible view of the elements in the array

// Retries are a premium feature now! Textio's free users only get 1 retry message, while pro members get an unlimited amount.

// Complete the getMessageWithRetriesForPlan function. It takes a plan variable as input as well as an array of 3 messages. You've been provided with constants representing the plan types at the top of the file.

// If the plan is a pro plan, return all the strings from the messages input in a slice.
// If the plan is a free plan, return the first 2 strings from the messages input in a slice.
// If the plan isn't either of those, return a nil slice and an error that says unsupported plan.
package main

import "errors"

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	// ?
	switch plan {
	case planFree:
		return messages[:2], nil
	case planPro:
		return messages[:], nil
	default:
		return nil, errors.New("unsupported plan")
	}
}
