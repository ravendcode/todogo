package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/mux"
)

// Todo model
type Todo struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title" gorm:"type:varchar(100);not null"`
	IsComplete bool      `json:"isComplete" gorm:"default:'false'"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// Validate method
func (t Todo) Validate(locale *Locale) error {
	return validation.ValidateStruct(&t,
		validation.Field(
			&t.Title,
			validation.Required.Error(ValidateTrans(locale, "validate.required", "model.title")),
			// validation.Required.Error(locale.T("validation.required", map[string]interface{}{"Field": locale.T("model.title")})),
			// validation.Required,
			validation.Length(2, 90).Error(ValidateMinMaxTrans(locale, "validate.length", "model.title", 2, 90)),
			// validation.Length(2, 90),
		),
		validation.Field(
			&t.IsComplete,
			// validation.Required.Error("не может быть пустым"),
			// validation.Required,
		),
	)
}

// Todos type
type Todos []Todo

// TodosResponse struct
type TodosResponse struct {
	Todos *Todos `json:"todos"`
}

// TodoResponse struct
type TodoResponse struct {
	Todo *Todo `json:"todo"`
}

// TodoHandler struct
type TodoHandler struct {
}

var todoRepo = new(TodoRepo)

// List method
func (t TodoHandler) List(w http.ResponseWriter, r *http.Request) {
	// render := RenderCtx(r.Context())
	// db := DbCtx(r.Context())

	render.JSON(w, TodosResponse{todoRepo.List()})
}

// Create method
func (t TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	todo := new(Todo)
	json.NewDecoder(r.Body).Decode(todo)
	if err := todo.Validate(LocaleCtx(r.Context())); err != nil {
		render.Status(http.StatusBadRequest).JSON(w, NewErrorValidate(err))
		return
	}
	todoRepo.Create(todo)

	render.Status(http.StatusCreated).JSON(w, TodoResponse{todo})
}

// Find method
func (t TodoHandler) Find(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	todo := new(Todo)
	todoRepo.Find(todo, id)
	if todo.ID == 0 {
		render.Status(http.StatusNotFound).JSON(w, NewError(LocaleCtx(r.Context()).T("error.todo_not_found")))
		return
	}
	render.JSON(w, TodoResponse{todo})
}

// Update method
func (t TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	todo := new(Todo)
	todoRepo.Find(todo, id)
	if todo.ID == 0 {
		render.Status(http.StatusNotFound).JSON(w, NewError(LocaleCtx(r.Context()).T("error.todo_not_found")))
		return
	}
	json.NewDecoder(r.Body).Decode(&todo)
	if err := todo.Validate(LocaleCtx(r.Context())); err != nil {
		render.Status(http.StatusBadRequest).JSON(w, NewErrorValidate(err))
		return
	}
	todoRepo.Save(todo)
	render.JSON(w, TodoResponse{todo})
}

// Destroy method
func (t TodoHandler) Destroy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	todo := new(Todo)
	todoRepo.Find(todo, id)
	if todo.ID == 0 {
		render.Status(http.StatusNotFound).JSON(w, NewError(LocaleCtx(r.Context()).T("error.todo_not_found")))
		return
	}
	todoRepo.Destroy(todo)
	render.SendStatus(w, http.StatusOK)
}
