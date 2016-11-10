package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello")
}

func foo(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Foo Dog")
}
