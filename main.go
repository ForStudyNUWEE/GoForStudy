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
	tmpl, _ := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	developer := Developer{"Бойчура М.В.", "НУВГП", 2023, "https://wiki.nuwm.edu.ua/index.php/Бойчура_Михайло_Володимирович"}

	tmpl.ExecuteTemplate(w, "index", developer)
}

func contacts(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/contacts.html")

	developer := Developer{"Бойчура М.В.", "НУВГП", 2023, "https://wiki.nuwm.edu.ua/index.php/Бойчура_Михайло_Володимирович"}

	tmpl.Execute(w, developer)
}

func login(w http.ResponseWriter, r *http.Request) {

				fmt.Print(r.PostForm, r.FormValue("username"), r.FormValue("password"))

		tmpl, _ := template.ParseFiles("templates/login.html", "templates/header.html", "templates/footer.html")

		developer := Developer{"Бойчура М.В.", "НУВГП", 2023, "https://wiki.nuwm.edu.ua/index.php/Бойчура_Михайло_Володимирович"}

		tmpl.ExecuteTemplate(w, "login", developer)
}

func main() {
	//forExample()
	http.HandleFunc("/", index)
	http.HandleFunc("/contacts/", contacts)
	http.HandleFunc("/login/", login)
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
