// A closure is a function that references variables from outside its own function body. The function may access and assign to the referenced variables.

package main

func adder() func(int) int {
	sum := 0
	return func(input int) int {
		sum += input
		return sum
	}
}
