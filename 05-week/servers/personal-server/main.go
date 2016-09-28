package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

var tpl *template.Template

type handler int

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	var h handler
	http.ListenAndServe(":8080", h)
}

func (m handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch req.URL.Path {
	case "/":
		err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", nil)
		if err != nil {
			log.Fatalln(err)
		}
	case "/cat":
		err := tpl.ExecuteTemplate(os.Stdout, "cat.gohtml", nil)
		if err != nil {
			log.Fatalln(err)
		}
	case "/dog":
		err := tpl.ExecuteTemplate(os.Stdout, "dog.gohtml", nil)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
