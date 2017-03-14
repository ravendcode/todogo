package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ravendcode/todong"
	"github.com/urfave/negroni"
)

func main() {
	config := todong.GetConfig()

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// http.ServeFile(w, r, "public/index.html")
		todong.RenderCtx(r.Context()).HTML(w, "index", nil)
	})
	// r.HandleFunc("/500", func(w http.ResponseWriter, r *http.Request) {
	// 	http.Error(w, "500 error", http.StatusInternalServerError)
	// })
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// http.ServeFile(w, r, "public/404.html")
		todong.RenderCtx(r.Context()).HTML(w, "404", nil)
		// todong.RenderCtx(r.Context()).HTML(w, "index", nil)
	})
	r.PathPrefix("/node_modules/").Handler(http.StripPrefix("/node_modules/", http.FileServer(http.Dir("node_modules"))))

	// r.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "public/assets/icons/favicon.ico")
	// })

	todong.APIRouter(r)

	n := negroni.Classic()
	// n := negroni.New()
	// n.Use(negroni.NewLogger())

	n.Use(negroni.HandlerFunc(todong.RenderMdw))

	n.UseHandler(r)

	// http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	// http.Handle("/public/node_modules/", http.StripPrefix("/public/node_modules/", http.FileServer(http.Dir("node_modules"))))
	// http.Handle("/public/js/", http.StripPrefix("/public/js/", http.FileServer(http.Dir("public/js"))))
	// http.Handle("/", http.FileServer(http.Dir("public")))
	// http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "public/assets/icons/favicon.ico")
	// })

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
