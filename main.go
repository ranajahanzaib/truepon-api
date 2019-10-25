package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	appengine.Main()
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Truepon Homepage.")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Find about Truepon here.")
}
