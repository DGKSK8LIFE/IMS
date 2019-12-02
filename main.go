package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		loginSite.ExecuteTemplate(w, "login.html", nil)
	})
	http.HandleFunc("/create.html", func(w http.ResponseWriter, r *http.Request) {
		createSite.ExecuteTemplate(w, "create.html", nil)
	})
	http.HandleFunc("/login", userAuth)
	http.HandleFunc("/create", createAccount)
	http.ListenAndServe(":8000", nil)
}

func userAuth(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	if len(email) > 0 && len(password) > 0 {
		db, err := sql.Open("sqlite3", "accounts.sqlite")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		rowExists(email, password, db)

		/* here we'll query the user account database to see if the this uname/password
		pair is a valid row */
	} // else {output some kind of message saying to fill out all forms}
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("username")
	password := r.FormValue("password")
	confirm := r.FormValue("confirm")
	if len(email) > 0 && len(password) > 0 && len(confirm) > 0 && password == confirm {
		/* here we check if the account already exists, if so, return an error message
		if not, write the row to the account database
		*/
	}
}

func rowExists(email, password string, db *sql.DB) bool {
	var exists bool
	query := fmt.Sprintf("SELECT * FROM ACCOUNTS WHERE email='%s' AND password='%s'", email, password)
	if err := db.QueryRow(query); err != nil {
		exists = false
	}
	exists = true
	return exists
}
