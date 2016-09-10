package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}


type item struct {
	Name, Price, Origin string
}

type meal struct {
	Time  string
	Items []item
}

type menu []meal

func main() {

	resturantOne := menu{
		meal{
			Time: "Breakfast",
			Items: []item{
				item{
					Name:   "Pankakes",
					Price:  "$15",
					Origin: "California",
				},
			},
		},
		meal{

			Time: "Lunch",
			Items: []item{
				item{
					Name:   "Hamburger",
					Price:  "$10",
					Origin: "Wyoming",
				},
				item{
					Name:   "Spaghetti",
					Price:  "$12",
					Origin: "China",
				},
			},
		},
		meal{

			Time: "Dinner",
			Items: []item{
				item{
					Name:   "Grilled Salmon",
					Price:  "$45",
					Origin: "Alaska",
				},
			},
		},
	}

	resturantTwo := menu{
		meal{
			Time: "Breakfast",
			Items: []item{
				item{
					Name:   "Pankakes",
					Price:  "$15",
					Origin: "California",
				},
			},
		},
		meal{

			Time: "Lunch",
			Items: []item{
				item{
					Name:   "Salad",
					Price:  "$10",
					Origin: "Wyoming",
				},
				item{
					Name:   "Spaghetti",
					Price:  "$12",
					Origin: "China",
				},
			},
		},
		meal{

			Time: "Dinner",
			Items: []item{
				item{
					Name:   "Grilled Salmon",
					Price:  "$45",
					Origin: "Alaska",
				},
			},
		},
	}

	resturants := []menu{resturantOne,resturantTwo}

	err := tpl.Execute(os.Stdout, resturants)
	if err != nil {
		log.Fatalln(err)
	}
}
