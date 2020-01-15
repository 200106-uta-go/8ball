package main

import (
	"fmt"
	"net/http"

	"github.com/200106-uta-go/8ball/config"
	"github.com/200106-uta-go/8ball/shaker"
)

func main() {
	println("Server is running on port 8080")

	http.Handle("/", http.FileServer(http.Dir("web")))
	http.HandleFunc("/8ball", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, config.Answers[shaker.Random(config.Abs, config.Answers)])
	})
	http.ListenAndServe(":8080", nil)
}
