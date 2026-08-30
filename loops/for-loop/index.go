// basic syntax for a for loop is:
// for INITIAL; CONDITION; AFTER{
//    do something
// }

// INITIAL is run once at the beginning of the loop and can create variables within the scope of the loop.
// CONDITION is checked before each iteration. If the condition doesn't pass then the loop breaks.
// AFTER is run after each iteration.

// At Textio we have a dynamic formula for determining how much a batch of bulk messages costs to send. Complete the bulkSend() function.

// It should return the total cost (as a float64) to send a batch of numMessages messages. Each message costs 1.0, plus an additional fee. The fee structure is:

// 1st message: 1.0 + 0.00
// 2nd message: 1.0 + 0.01
// 3rd message: 1.0 + 0.02
// 4th message: 1.0 + 0.03
// ...
package main

func bulkSend(numMessages int) float64 {
	// ?
	totalCost := 0.0
	for i := 1; i <= numMessages; i++ {
		totalCost += 1.0 + float64(i-1)*0.01
	}
	return totalCost
}
