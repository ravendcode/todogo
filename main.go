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

func init() {
}

func main() {
	fmt.Println("ENV is", config.ENV)

	var err error
	db, err = gorm.Open("sqlite3", config.DbPath)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.LogMode(config.ENV == "development")
	MigrateDb(db)

	r := mux.NewRouter()

	IndexRoutes(r)
	APIRoutes(r)

	n := negroni.Classic()

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
