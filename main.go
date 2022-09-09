package main // import "hello-render"

import (
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	http.HandleFunc("/healthz", health)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":5525", nil)
}
