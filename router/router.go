package router

import (
	"net/http"

	"github.com/DcWire/path-finder-go/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api", controller.PassHtml).Methods("GET")
	router.HandleFunc("/api/init-matrix", controller.InitHtml).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static/")))
	return router
}
