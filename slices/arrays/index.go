// Complete the getMessageWithRetries function. It takes three strings and returns:

// An array of 3 strings
// An array of 3 integers
// The returned string array contains the original messages. The first is the primary message, the second is the first reminder, and the third is the last reminder.

// The integers in the integer array represent the cost of sending each message. The cost of each message is equal to the length of the message, plus the length of any previous messages. For example:

// "hello" costs 5
// "world" costs 5, adding "hello" makes total cost 10 (5 + 5)
// "!" costs 1, adding the previous messages makes total cost 11 (5 + 5 + 1)

package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	// ?
	messages := [3]string{primary, secondary, tertiary}
	messageCost := len(primary)
	firstReminderCost := messageCost + len(secondary)
	secondReminderCost := firstReminderCost + len(tertiary)

	costArray := [3]int{messageCost, firstReminderCost, secondReminderCost}

	return messages, costArray
}
