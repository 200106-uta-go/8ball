// Package shaker is an 8ball library
package shaker

import "math/rand"

// Random is a function that takes in a bool and a string slice
// which returns a random number. To use properly, please
// seed your random.
func Random(abs bool, answers []string) int {
	var n int
	if abs {
		n = 2
	} else {
		n = len(answers)
	}
	return rand.Intn(n)
}
