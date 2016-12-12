package main

import "net/http"

func getSession(w http.ResponseWriter, req *http.Request) user {
	// your code here
}

func alreadyLoggedIn(req *http.Request) bool {
	// your code here
}

func protected(req *http.Request) (user, error) {
	var u user
	c, err := req.Cookie("session")
	if err != nil {
		return u, err
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		return u, err
	}
	u = dbUsers[un]
	return u, nil
}