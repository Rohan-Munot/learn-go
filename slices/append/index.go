// Complete the getDayCosts() function. It accepts a slice of cost structs and a day int, and it returns a float64 slice containing that day's costs.

// Create an empty, non-nil float64 slice.
// Use append() to add each cost's value when its day matches the day argument.
// Return the slice.

package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	dayCosts := make([]float64, 0, len(costs))
	for _, cost := range costs {
		if cost.day == day {
			dayCosts = append(dayCosts, cost.value)
		}
	}
	return dayCosts
}
