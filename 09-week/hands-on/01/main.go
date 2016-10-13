package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func foo(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "foo.gohtml", req.Method)
}

func bar(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "bar.gohtml", req.URL.Path)
}

func me(res http.ResponseWriter, _ *http.Request) {
	d := struct {
		FName string
		LName string
	}{
		"Jeremy",
		"Perkins",
	}

	tpl.ExecuteTemplate(res, "me.gohtml", d)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", foo)
	mux.HandleFunc("/bar/", bar)
	mux.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", mux)
}
