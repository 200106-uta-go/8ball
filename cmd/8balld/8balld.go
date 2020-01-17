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
		var name = r.FormValue("name")
		if name == "" {
			name = "Anonymous"
		}
		fmt.Fprint(w, config.Answers[shaker.Random(config.Abs, config.Answers)], " ", name)
	})
	http.ListenAndServe(":8080", nil)
}
