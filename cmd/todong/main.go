package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ravendcode/todong"
	"github.com/urfave/negroni"
)

func main() {
	config := todong.GetConfig()

	// r := todong.Router()
	r := mux.NewRouter()
	todong.APIRouter(r)
	http.Handle("/api/", r)

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	n := negroni.Classic()
	// n := negroni.New()
	// n.Use(negroni.HandlerFunc(middlewares.ContextDbMiddleware))
	// n.Use(negroni.HandlerFunc(middlewares.ContextViewMiddleware))
	n.Use(negroni.HandlerFunc(todong.RenderMdw))

	n.UseHandler(r)

	http.Handle("/node_modules/", http.StripPrefix("/node_modules/", http.FileServer(http.Dir("node_modules"))))
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/assets/icons/favicon.ico")
	})
	// http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Printf("Server listening on localhost:%s\n", config.Port)

	s := &http.Server{
		Addr:           ":" + config.Port,
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
