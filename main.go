package main

import (
	"fmt"
	"os"

	"github.com/200106-uta-go/8ball/shaker"
)

func main() {
	user := os.Getenv("USER")
	fmt.Println(answers[shaker.Random(abs, answers)], user)

}
