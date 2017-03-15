package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/urfave/negroni"
)

var config = NewConfig()

var db *gorm.DB

func main() {
	fmt.Println("Env is", config.Env)

	var err error
	db, err = gorm.Open("sqlite3", config.DbPath)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.LogMode(config.Env == "development")
	MigrateDb(db)

	r := mux.NewRouter()

	IndexRoutes(r)
	APIRoutes(r)

	n := negroni.Classic()
	// n := negroni.New()
	// n.Use(negroni.NewStatic(http.Dir("public")))
	// n.Use(negroni.NewLogger())
	// n.Use(negroni.NewRecovery())
	n.Use(negroni.HandlerFunc(LocaleMdw))
	n.Use(negroni.HandlerFunc(RenderMdw))
	n.Use(negroni.HandlerFunc(DbMdw(db)))

	n.UseHandler(r)

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
