package main

import (
	"io"
	"net/http"
)

func i(res http.ResponseWriter, _ *http.Request) {
	io.WriteString(res, "index index index")
}

func d(res http.ResponseWriter, _ *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, _ *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	http.HandleFunc("/", i)
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}
