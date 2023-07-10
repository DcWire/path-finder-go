package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

// func PassCSS(w http.ResponseWriter, r *http.Request) {

// }
func PassHtml(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got request")
	person := Person{
		"Anas",
		22,
	}

	tmp, err := template.ParseFiles("views/home.html")
	if err != nil {
		panic(err)
	}

	err = tmp.Execute(w, person)

	if err != nil {
		panic(err)
	}

}
