package todong

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
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	IsComplete bool      `json:"isComplete"`
	CreatedAt  time.Time `json:"createdAt"`
}

// Validate method
func (t Todo) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(
			&t.Title,
			// validation.Required.Error("не может быть пустым"),
			validation.Required,
			validation.Length(2, 50),
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

func (t TodoHandler) list(w http.ResponseWriter, r *http.Request) {
	render := RenderCtx(r.Context())
	render.JSON(w, TodosResponse{&todoList})
}

func (t TodoHandler) create(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	if err := todo.Validate(); err != nil {
		render.Status(http.StatusBadRequest).JSON(w, NewErrorValidate(err))
		return
	}
	todoCreateRepo(&todo)

	render.Status(http.StatusCreated).JSON(w, TodoResponse{&todo})
}

func (t TodoHandler) find(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	index, todo := todoFindRepo(id)
	if index == -1 {
		render.Status(http.StatusNotFound).JSON(w, NewError("Todo Not Found"))
		return
	}
	render.JSON(w, TodoResponse{&todo})
}

func (t TodoHandler) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	index, todo := todoFindRepo(id)
	if index == -1 {
		render.Status(http.StatusNotFound).JSON(w, NewError("Todo Not Found"))
		return
	}
	json.NewDecoder(r.Body).Decode(&todo)
	if err := todo.Validate(); err != nil {
		render.Status(http.StatusBadRequest).JSON(w, NewErrorValidate(err))
		return
	}
	todoList[index] = todo
	render.JSON(w, TodoResponse{&todo})
}

func (t TodoHandler) destroy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	index, _ := todoFindRepo(id)
	if index == -1 {
		render.Status(http.StatusNotFound).JSON(w, NewError("Todo Not Found"))
		return
	}
	todoDestroyRepo(index)
	render.SendStatus(w, http.StatusOK)
}
