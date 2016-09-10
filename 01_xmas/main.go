package main

import (
	"fmt"
	"text/template"
	"os"
)


func main() {

	tpl, err := template.ParseFiles("letter.txt")
	if err != nil {
		fmt.Println("There was an error parsing the file", err)
	}

	friends := []string{"Alex", "Connor", "Bob", "John", "Todd"}

	err = tpl.Execute(os.Stdout, friends)
	if err != nil {
		fmt.Println("There was an error executing the template", err)
	}
}
