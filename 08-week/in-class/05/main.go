package main

import (
	"io"
	"net/http"
)

type hotdog int
func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int
func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

type hamster int

func (h hamster) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "hamster hamster hamster")
}

func main() {
	var d hotdog
	var c hotcat
	var h hamster

	http.Handle("/dog", d)
	http.Handle("/cat", c)
	http.Handle("/hamster", h)

	http.ListenAndServe(":8080", nil)
}

// change code to use DefaultServeMux
// add a route for hamsters