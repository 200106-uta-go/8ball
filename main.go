package main

import (
	"flag"
	"math/rand"
	"time"
)

var answers [8]string
var abs bool

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

	flag.BoolVar(&abs, "abs", false, "absolute answers only")
	flag.Parse()
}

func random() int {
	var n int
	if abs {
		n = 2
	} else {
		n = len(answers)
	}
	return rand.Intn(n)
}

func main() {
	println(answers[random()])
}
