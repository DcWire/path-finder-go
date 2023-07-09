package router

import (
	"github.com/DcWire/path-finder-go/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api", controller.PassHtml).Methods("GET")
	return router
}
