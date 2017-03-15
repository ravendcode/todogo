package main

import (
	"context"
	"net/http"

	"github.com/jinzhu/gorm"
)

type contextKey string

func (c contextKey) String() string {
	return "main context key " + string(c)
}

var contextKeyRender = contextKey("render")
var contextKeyDb = contextKey("db")

// RenderMdw func
func RenderMdw(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := context.WithValue(r.Context(), contextKeyRender, NewRender())
	next(w, r.WithContext(ctx))
}

// DbMdw func
func DbMdw(db *gorm.DB) func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		ctx := context.WithValue(r.Context(), contextKeyDb, db)
		next(w, r.WithContext(ctx))
	}
}
