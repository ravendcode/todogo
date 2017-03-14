package todong

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

// Render struct
type Render struct {
	layout string
	status int
}

// HTML func
func (r *Render) HTML(w io.Writer, name string, data interface{}) error {
	output, err := template.New("").Delims("[[", "]]").ParseFiles(
		fmt.Sprintf("templates/%s.html", name),
		fmt.Sprintf("templates/layouts/%s.html", r.layout),
		// fmt.Sprintf("templates/partials/_nav.html"),
		// fmt.Sprintf("templates/partials/_user.html"),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	templates := template.Must(output, err)
	return templates.ExecuteTemplate(w, "base", data)
}

// Layout method
func (r *Render) Layout(name string) *Render {
	r.layout = name
	return r
}

// Status method
func (r *Render) Status(status int) *Render {
	r.status = status
	return r
}

// SendStatus method
func (r *Render) SendStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(r.status)
}

// JSON method
func (r *Render) JSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(r.status)
	r.status = http.StatusOK
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// NewRender create new Render instance
func NewRender() *Render {
	return &Render{layout: "base", status: http.StatusOK}
}

var render = NewRender()
