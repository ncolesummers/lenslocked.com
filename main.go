package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Welcome to my awesome site!<h1>") // First argument for Fprint is where you are writing
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}