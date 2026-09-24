// Variadic functions are functions that can take a variable number of arguments of the same type.
// variadic argumemnt must be the last argument in the function signature.

// We need to sum up the costs of all individual messages so we can send an end-of-month bill to our customers.

// Complete the sum function to return the sum of all inputs.

// Take note of how the variadic inputs and the spread operator are used in the test suite.

package main

func sum(nums ...int) int {
	total := 0
	for i := range nums {
		total += nums[i]
	}
	return total
}
