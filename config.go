package main

import (
	"flag"
	"math/rand"
	"time"
)

var answers []string
var abs bool

func init() {
	answers = []string{
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
