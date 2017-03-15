package main

import (
	"context"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/nicksnyder/go-i18n/i18n"
)

type contextKey string

func (c contextKey) String() string {
	return "main context key " + string(c)
}

var contextKeyRender = contextKey("render")
var contextKeyDb = contextKey("db")
var contextKeyLocale = contextKey("locale")

// Locale struct
type Locale struct {
	Locale string
	T      i18n.TranslateFunc
}

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

// LocaleMdw func
func LocaleMdw(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	locale := new(Locale)

	if value, ok := r.URL.Query()["locale"]; ok {
		localeFind := false
		for _, lang := range config.Locales {
			if value[0] == lang {
				locale.Locale = value[0]
				localeFind = true
			}
		}
		if !localeFind {
			locale.Locale = config.Locales[0]
		}
	} else {
		if value, ok := r.URL.Query()["lang"]; ok {
			localeFind := false
			for _, lang := range config.Locales {
				if value[0] == lang {
					locale.Locale = value[0]
					localeFind = true
				}
			}
			if !localeFind {
				locale.Locale = config.Locales[0]
			}
		} else {
			locale.Locale = config.Locales[0]
		}
	}

	T, _ := i18n.Tfunc(locale.Locale)
	locale.T = T
	ctx := context.WithValue(r.Context(), contextKeyLocale, locale)
	next(w, r.WithContext(ctx))
}
