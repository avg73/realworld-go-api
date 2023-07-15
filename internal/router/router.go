package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigureRouter(r *mux.Router) error {
	r.HandleFunc("/", HomeHandler)
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", ApiHendler)
	return nil
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func ApiHendler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello api")
}
