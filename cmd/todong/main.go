package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/ravendcode/todong"
)

func main() {
	config := todong.GetConfig()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	// n := negroni.Classic()
	// n.Use(negroni.HandlerFunc(middlewares.ContextDbMiddleware))
	// n.Use(negroni.HandlerFunc(middlewares.ContextViewMiddleware))
	// n.UseHandler(router)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.Handle("/node_modules/", http.StripPrefix("/node_modules/", http.FileServer(http.Dir("node_modules"))))

	fmt.Printf("Server listening on localhost:%s\n", config.Port)

	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
