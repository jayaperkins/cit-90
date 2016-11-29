package main

import (
	"github.com/satori/go.uuid"
	"net/http"
	"html/template"
)

type user struct {
	First string
	Last string
}

var tpl *template.Template
var db = map[string]user{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", login)
	http.HandleFunc("/bar", user)
	http.ListenAndServe(":8080", nil)
}

func login(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		http.SetCookie(w, c)

	}

	var u user
	if req.Method == http.MethodPost {
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{f,l}
		db[c.Value] = u
	}

	tpl.ExecuteTemplate(w, "login.html", nil)


}

func user(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u, ok := db[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}


	tpl.ExecuteTemplate(w, "user.html", u)
}
