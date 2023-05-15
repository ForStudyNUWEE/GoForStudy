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

func contacts(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/contacts.html")

	developer := Developer{"Бойчура М.В.", "НУВГП", 2023, "https://wiki.nuwm.edu.ua/index.php/Бойчура_Михайло_Володимирович"}

	tmpl.Execute(w, developer)
}

func main() {
	//forExample()
	http.HandleFunc("/", index)
	http.HandleFunc("/contacts/", contacts)
	http.ListenAndServe(":8080", nil)
}

func forExample() {
	//var x int
	//var x int = 5
	//var x = 5
	//x := 5
	//fmt.Print(x)

	//var y Developer
	//var y Developer = Developer{"", "", 5, ""}
	//var y = Developer{"fsd", "ds", 5, "ds"}
	//y := Developer{"fsd", "ds", 5, "ds"}
	//fmt.Print(y)
	
}
