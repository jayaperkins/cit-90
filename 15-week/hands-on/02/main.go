package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("You are at foo - the request method was: ", req.Method)
	tpl.ExecuteTemplate(w, "foo.html", nil)
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("You are at bar - the request method was: ", req.Method)
	tpl.ExecuteTemplate(w, "bar.html", nil)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("You are at barred - the request method was: ", req.Method)
	tpl.ExecuteTemplate(w, "barred.html", nil)
}
