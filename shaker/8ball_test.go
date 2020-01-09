package shaker

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	answers := []string{"Yes", "No", "Maybe"}
	n := Random(false, answers)
	if n < 0 {
		t.Error("n should not be less than 0")
	} else if n >= len(answers) {
		t.Errorf("%d is greater than length of answers", n)
	}
}

func ExampleRandom() {
	answers := []string{"Yes", "No", "Maybe"}
	n := Random(false, answers)
	fmt.Println(n >= 0)
	// Output: true
}
