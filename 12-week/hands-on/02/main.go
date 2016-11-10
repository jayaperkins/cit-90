package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/snow", snow)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello")
}

func foo(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Foo Dog")
}

func snow(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "snow.jpeg")
}
