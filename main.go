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
	if n == 0 {
		println("No")
	} else if n == 1 {
		println("Yes")
	} else if n == 2 {
		println("Maybe")
	}
}
