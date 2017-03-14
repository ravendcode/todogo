package todong

import (
	"context"
	"net/http"
)

type contextKey string

func (c contextKey) String() string {
	return "todong context key " + string(c)
}

var contextKeyRender = contextKey("render")

// RenderMdw func
func RenderMdw(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := context.WithValue(r.Context(), contextKeyRender, NewRender())
	next(w, r.WithContext(ctx))
}
