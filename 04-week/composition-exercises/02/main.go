package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type hotel struct {
	Name, Address, City, Zip, Region string
}

type hotels []hotel

func main() {

	h := hotels{
		hotel{
			Name: "Hilton",
			Address: "111 West Oak Street",
			City: "Fresno",
			Zip: "43894",
			Region: "Southern",
		},

		hotel{
			Name: "Best Western",
			Address: "4839 North Pine Ave",
			City: "Modesto",
			Zip: "54892",
			Region: "Northern",
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}