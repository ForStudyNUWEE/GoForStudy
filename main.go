package main

import (
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

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
