package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Developer struct {
	Name         string
	Organization string
	Year         int
	Url          string
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")

	developer := Developer{"Бойчура М.В.", "НУВГП", 2023, "https://wiki.nuwm.edu.ua/index.php/Бойчура_Михайло_Володимирович"}

	tmpl.Execute(w, developer)
}

func contacts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "09888234685")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/contacts/", contacts)
	http.ListenAndServe(":8080", nil)
}
