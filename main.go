package main

import (
	"math/rand"
	"time"
)

func main() {
	var t int64
	t = int64(time.Now().Second())
	rand.Seed(t)
	n := rand.Intn(3)

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

	println(answers[n])
}
