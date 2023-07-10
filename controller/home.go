package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Maze struct {
	Maze [][]int
}

type Size struct {
	Row json.Number `json:"row"`
	Col json.Number `json:"col"`
}

// func PassCSS(w http.ResponseWriter, r *http.Request) {

// }
var maze *Maze

func InitHtml(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Body)

	err := json.NewDecoder(r.Body).Decode(&maze)
	if err != nil {
		panic(err)
	}
	fmt.Println(maze.Maze)

}
func PassHtml(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got request")

	tmp, err := template.ParseFiles("views/Home/home.html")
	if err != nil {
		panic(err)
	}

	err = tmp.Execute(w, nil)

	if err != nil {
		panic(err)
	}

}
