package main

import "net/http"

import "fmt"

import "io/ioutil"

func main() {
	resp, _ := http.Get("http://localhost:8080/8ball")
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(body))
}
