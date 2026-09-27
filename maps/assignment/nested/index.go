// Complete the getNameCounts function. It takes a slice of strings names and returns a nested map. The parent map's keys are all the unique first characters (see runes) of the names, the nested maps keys are all the names themselves, and the value is the count of each name.

package main

func getNameCounts(names []string) map[rune]map[string]int {
	// Your code here
	nameCounts := make(map[rune]map[string]int)
	for _, name := range names {
		initial := []rune(name)[0]
		if nameCounts[initial] == nil {
			nameCounts[initial] = make(map[string]int)
		}
		nameCounts[initial][name]++
	}
	return nameCounts
}
