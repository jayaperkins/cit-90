package main

import (
	"net/http"
	"html/template"
	"log"
	"net/url"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
	}

	data := struct{
		Method string
		Submissions url.Values
		URL *url.URL
	}{
		req.Method,
		req.Form,
		req.URL,
	}
	tpl.ExecuteTemplate(res, "foo.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("foo.gohtml"))
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}

