package main

import (
	"io"
	"net/http"
)


type hampster int

func (h hampster) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "<h1>Hampster</h1>")
}

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "<h1>Dog</h1>")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "<h1>Cat</h1>")
}

func main() {

	var h hampster
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/hampster", h)
	mux.Handle("/dog", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)

}
