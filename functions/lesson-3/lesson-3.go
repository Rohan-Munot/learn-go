package main

import "fmt"

func getMonthlyPrice(tier string) int {
	switch tier {
	case "basic":
		return 10000
	case "premium":
		return 15000
	case "enterprise":
		return 50000
	default:
		return 0
	}
}

// don't touch below this line

func main() {
	test("basic")
	test("premium")
	test("enterprise")
	test("invalid")
	test("")
}

func test(tier string) {
	fmt.Printf("tier=%q -> %d\n", tier, getMonthlyPrice(tier))
}
