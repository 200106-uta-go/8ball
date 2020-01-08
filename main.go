package main

import (
	"math/rand"
	"time"
)

func main() {
	answers := [8]string{
		"No",
		"Yes",
		"Maybe",
		"Try Again",
		"Pray",
		"I don't know",
		"It's unclear",
		"Can you repeat the question?",
	}

	var t int64
	t = time.Now().UnixNano()
	rand.Seed(t)
	n := rand.Intn(len(answers))

	println(answers[n])
}
