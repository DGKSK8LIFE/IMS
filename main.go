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
		if row := rowExists(email, password, db); row == true {
			fmt.Fprint(w, "<h1 style='text-align: center;'>Welcome!</h1>")
		} else if row == false {
			// deny access
			loginSite.ExecuteTemplate(w, "login.html", nil)
		}

	} else {
		fmt.Fprint(w, "<h1 style='text-align: center;'>please fill out all forms</h1>")
	}
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	db, err := sql.Open("sqlite3", "accounts.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if row := rowExists(email, password, db); row == false {
		query := fmt.Sprintf("INSERT INTO accounts (email, password) \n VALUES ('%s', '%s');", email, password)
		db.Exec(query)
	} else if row == true {
		fmt.Fprint(w, "<h1 style='text-align: center;'>Account already exists!</h1>")
	}
}

func rowExists(email, password string, db *sql.DB) bool {
	var exists bool
	query := fmt.Sprintf("SELECT * FROM ACCOUNTS WHERE email='%s' AND password='%s'", email, password)
	if err := db.QueryRow(query).Scan(&email, &password); err != nil && err != sql.ErrNoRows {
		log.Fatal("database error, we're fucked")
	} else if err == sql.ErrNoRows {
		return false
	} else {
		exists = true
	}
	return exists
}

func accountTaken(email string, db *sql.DB) bool {
	var exists bool
	query := fmt.Sprintf("SELECT * FROM ACCOUNTS WHERE email='%s'", email)
	if err := db.QueryRow(query).Scan(&email); err != nil && err != sql.ErrNoRows {
		log.Fatal("database error, we're fucked")
	} else if err == sql.ErrNoRows {
		exists = false
	} else {
		exists = true
	}
	return exists
}
