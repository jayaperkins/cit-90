package main

import (
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbUsers["james@bond.com"] = user{"james@bond.com", "notstirred", "James", "Bond"} // populate our db
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}

	// if the user exists already, get user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, p, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.html", u)
}

func bar(w http.ResponseWriter, req *http.Request) {

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.html", u)
}

func login(w http.ResponseWriter, req *http.Request) {

	// is the user already logged in?
	c, _ := req.Cookie("session")
	if c != nil {
		un := dbSessions[c.Value]
		if _, ok := dbUsers[un]; ok {
			http.Redirect(w, req, "/", http.StatusSeeOther)
			return
		}
	}

	var u user
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		u = dbUsers[un]
		if u.Password != p {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un
	}

	tpl.ExecuteTemplate(w, "login.html", u)
}

func logout(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("session")
	if err != nil {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusSeeOther)
		return
	}

	delete(dbSessions, c.Value)

	c = &http.Cookie {
		Name: "session",
		Value: "",
		MaxAge: -1,
	}

	http.SetCookie(w,c)
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}