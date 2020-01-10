package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"os"
	"time"
)

var answers []string
var abs bool

const CONFIGFILE string = "conf.json"

type Configuration struct {
	File string `json:"stdanswers"`
}

func init() {
	config := Configuration{}
	c, _ := os.Open(CONFIGFILE)
	json.NewDecoder(c).Decode(&config)

	f, err := os.Open(config.File)
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
