package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// APIRoutes func
func APIRoutes(router *mux.Router) {
	r := router.StrictSlash(false).PathPrefix("/api").Subrouter()

	todoHandler := &TodoHandler{}
	r.HandleFunc("/todos", todoHandler.List).Methods("GET")
	r.HandleFunc("/todos", todoHandler.Create).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", todoHandler.Find).Methods("GET")
	r.HandleFunc("/todos/{id:[0-9]+}", todoHandler.Update).Methods("PATCH")
	r.HandleFunc("/todos/{id:[0-9]+}", todoHandler.Destroy).Methods("DELETE")
}

// IndexRoutes func
func IndexRoutes(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/assets/index.html")
	})

	router.PathPrefix("/node_modules/").Handler(
		http.StripPrefix("/node_modules/", http.FileServer(http.Dir("node_modules"))))

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// http.ServeFile(w, r, "public/404.html")
		http.ServeFile(w, r, "public/assets/index.html")
	})
}
