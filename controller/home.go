package controller

import (
	"html/template"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func PassHtml(w http.ResponseWriter, r *http.Request) {
	person := Person{
		"Anas",
		22,
	}

	tmp, err := template.ParseFiles("views/Home/home.html")
	if err != nil {
		panic(err)
	}

	err = tmp.Execute(w, person)

	if err != nil {
		panic(err)
	}

}
