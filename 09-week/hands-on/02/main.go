package main

import (
	"html/template"
	"net/http"
	"net/url"
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

func me(res http.ResponseWriter, req *http.Request) {
	d := struct {
		FName string
		LName string
		URL   *url.URL
	}{
		"Jeremy",
		"Perkins",
		req.URL,
	}

	tpl.ExecuteTemplate(res, "me.gohtml", d)
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar/", bar)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
