package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})

	return route
}
