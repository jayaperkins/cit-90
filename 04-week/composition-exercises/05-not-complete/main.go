package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func parse() {

	records := prs("table.csv")

	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(res, records)
	if err != nil {
		log.Fatalln(err)
	}
}
