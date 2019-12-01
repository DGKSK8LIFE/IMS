package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	loginSite  *template.Template
	createSite *template.Template
)

func init() {
	loginSite = template.Must(template.ParseGlob("login.html"))
	createSite = template.Must(template.ParseGlob("create.html"))
}

func main() {
	http.HandleFunc("/", handleLoginSite)
	http.HandleFunc("/create.html", handleCreateSite)
	http.HandleFunc("/login", userAuth)
	http.HandleFunc("/create", createAccount)
	http.ListenAndServe(":8000", nil)
}

func handleLoginSite(w http.ResponseWriter, r *http.Request) {
	loginSite.ExecuteTemplate(w, "login.html", nil)
}

func handleCreateSite(w http.ResponseWriter, r *http.Request) {
	createSite.ExecuteTemplate(w, "create.html", nil)
}

func userAuth(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)
	if len(username) > 0 && len(password) > 0 {
		/* here we'll query the user account database to see if the this uname/password
		pair is a valid row */
	} // else {output some kind of message saying to fill out all forms}
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	confirm := r.FormValue("confirm")
	if len(username) > 0 && len(password) > 0 && len(confirm) > 0 && password == confirm {
		/* here we check if the account already exists, if so, return an error message
		if not, write the row to the account database
		*/
	}
}
