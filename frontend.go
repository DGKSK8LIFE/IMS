package main

import (
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
	
	http.ListenAndServe(":8000", nil)
}

func handleLoginSite(w http.ResponseWriter, r *http.Request) {
	loginSite.Execute(w, "login.html")
}
