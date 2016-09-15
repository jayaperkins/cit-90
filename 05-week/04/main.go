package main

import (
	"io"
	"net/http"
	"html/template"
	"log"
)

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func bar(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bar ran")
}

func name(res http.ResponseWriter, req *http.Request) {
	//io.WriteString(res, "Jeremy Perkins")
	tpl, err := template.ParseFiles("tpl.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/name", name)
	http.ListenAndServe(":8080", nil)
}

