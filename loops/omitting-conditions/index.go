// Complete the maxMessages function. The parameter thresh is your total budget in pennies. Return the maximum number of whole messages you can send without the total cost ever exceeding thresh.

// Each message costs 100 pennies, plus an additional fee. The fee structure is:

// 1st message: 100 + 0
// 2nd message: 100 + 1
// 3rd message: 100 + 2
// 4th message: 100 + 3

package main

func maxMessages(thresh int) int {
	// ?
	totalCost := 0
	for i := 1; ; i++ {
		messageCost := 100 + (i-1)*1
		totalCost += messageCost
		if totalCost > thresh {
			return i - 1
		}
	}
}
