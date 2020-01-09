package main

import (
	"bufio"
	"flag"
	"log"
	"math/rand"
	"os"
	"time"
)

var answers []string
var abs bool

func init() {
	f, err := os.Open("config.txt")
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewScanner(f)

	for reader.Scan() {
		answers = append(answers, reader.Text())
	}

	rand.Seed(time.Now().UnixNano())

	flag.BoolVar(&abs, "abs", false, "absolute answers only")
	flag.Parse()
}
