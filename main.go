package main

import (
	"math/rand"
	"time"
)

var answers [8]string

func init() {
	answers = [8]string{
		"No",
		"Yes",
		"Maybe",
		"Try Again",
		"Pray",
		"I don't know",
		"It's unclear",
		"Can you repeat the question?",
	}

	rand.Seed(time.Now().UnixNano())
}

func random() int {
	return rand.Intn(len(answers))
}

func main() {
	println(answers[random()])
}
