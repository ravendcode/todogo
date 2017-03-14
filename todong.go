package todong

import (
	"github.com/gorilla/mux"
)

// APIRouter func
func APIRouter(router *mux.Router) {

	r := router.StrictSlash(false).PathPrefix("/api").Subrouter()

	todoHandler := TodoHandler{}
	r.HandleFunc("/todos", todoHandler.list).Methods("GET")
	r.HandleFunc("/todos", todoHandler.create).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", todoHandler.find).Methods("GET")
	r.HandleFunc("/todos/{id:[0-9]+}", todoHandler.update).Methods("PATCH")
	r.HandleFunc("/todos/{id:[0-9]+}", todoHandler.destroy).Methods("DELETE")
}
