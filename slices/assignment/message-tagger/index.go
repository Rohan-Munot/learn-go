// Complete the tagMessages function. It should take a slice of sms messages, and a function (that takes an sms as input and returns a slice of strings) as inputs. And it should return a slice of sms messages.

// It should loop through each message and set the tags to the result of the passed in function.
// Be sure to modify the messages of the original slice using bracket notation messages[i].

// Complete the tagger function. It should take an sms message and return a slice of strings.

// Return an initialized slice, even if no tags match. No nil slices.
// For any message that contains "urgent" (regardless of casing) in the content, the Urgent tag should be applied first.
// For any message that contains "sale" (regardless of casing), the Promo tag should be applied second.

package main

import "strings"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	// ?
	for i, msg := range messages {
		messages[i].tags = tagger(msg)
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	if strings.Contains(strings.ToLower(msg.content), "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(strings.ToLower(msg.content), "sale") {
		tags = append(tags, "Promo")
	}
	return tags
}
