package config

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"os"
	"time"
)

var Answers []string
var Abs bool
var Http bool
var Port int

const CONFIGFILE string = "conf.json"

type Configuration struct {
	File string `json:"stdanswers"`
	Port int    `json:"port"`
}

func init() {
	config := Configuration{}
	c, _ := os.Open(CONFIGFILE)
	json.NewDecoder(c).Decode(&config)
	Port = config.Port

	f, err := os.Open(config.File)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewScanner(f)

	for reader.Scan() {
		Answers = append(Answers, reader.Text())
	}

	rand.Seed(time.Now().UnixNano())

	flag.BoolVar(&Abs, "abs", false, "absolute answers only")
	flag.BoolVar(&Http, "http", false, "start http server")
	flag.Parse()
}
