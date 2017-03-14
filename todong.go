package todong

import (
	"github.com/gorilla/mux"
)

// APIRouter func
func APIRouter(router *mux.Router) {

	r := router.StrictSlash(false).PathPrefix("/api").Subrouter()

	r.HandleFunc("/todos", todoListHandler).Methods("GET")
	r.HandleFunc("/todos", todoCreateHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", todoGetHandler).Methods("GET")
	r.HandleFunc("/todos/{id:[0-9]+}", todoUpdateHandler).Methods("PATCH")
	r.HandleFunc("/todos/{id:[0-9]+}", todoDeleteHandler).Methods("DELETE")
}
