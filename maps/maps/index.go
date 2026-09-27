// Complete the getUserMap function. It takes a slice of names and a slice of phone numbers, and returns a map of name -> user structs and an error. A user struct just contains a user's name and phone number. The first element in the names slice pairs with the first phone number, and so on.

// If the length of names and phoneNumbers is not equal, return an error with the string "invalid sizes".

package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	// ?
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	userMap := make(map[string]user)
	for i := range names {
		userMap[names[i]] = user{names[i], phoneNumbers[i]}
	}
	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}
