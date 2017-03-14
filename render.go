package todong

import (
	"encoding/json"
	"net/http"
)

// Render struct
type Render struct {
	Layout string
	Status int
}

// func (v *Render) HTML(w io.Writer, name string, data interface{}) error {
// 	output, err := template.ParseFiles(
// 		fmt.Sprintf("templates/%s.html", name),
// 		fmt.Sprintf("templates/layouts/%s.html", v.Layout),
// 		fmt.Sprintf("templates/partials/_nav.html"),
// 		fmt.Sprintf("templates/partials/_user.html"),
// 	)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	templates := template.Must(output, err)
// 	return templates.ExecuteTemplate(w, "base", data)
// }

// SetLayout method
func (v *Render) SetLayout(name string) *Render {
	v.Layout = name
	return v
}

// SetStatus method
func (v *Render) SetStatus(status int) *Render {
	v.Status = status
	return v
}

// SendStatus method
func (v *Render) SendStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(v.Status)
}

// JSON method
func (v *Render) JSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(v.Status)
	v.Status = http.StatusOK
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// NewRender create new Render instance
func NewRender() *Render {
	return &Render{Layout: "base", Status: http.StatusOK}
}
