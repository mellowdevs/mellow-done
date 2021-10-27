package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mellowdevs/mellow-done/controllers/auth"
)

func Init(db *sql.DB) *mux.Router {
	route := mux.NewRouter()
	auth_route := "/api/v1/auth"

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})

	route.HandleFunc(auth_route+"/register", func(rw http.ResponseWriter, r *http.Request) {
		auth.Register(rw, r, db)
	}).Methods("POST")
	route.HandleFunc(auth_route+"/login", func(rw http.ResponseWriter, r *http.Request) {
		auth.Login(rw, r, db)
	}).Methods("POST")
	route.HandleFunc(auth_route, func(rw http.ResponseWriter, r *http.Request) {
		auth.GetUser(rw, r)
	}).Methods("GET")
	route.HandleFunc(auth_route+"/refresh", func(rw http.ResponseWriter, r *http.Request) {
		auth.Refresh(rw, r)
	}).Methods("GET")
	route.HandleFunc(auth_route+"/logout", func(rw http.ResponseWriter, r *http.Request) {
		auth.Logout(rw, r)
	}).Methods("GET")
	return route
}
