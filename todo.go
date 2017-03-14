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

var todoList = Todos{
	{ID: 1, Title: "Create chat", IsComplete: false, CreatedAt: time.Now().Add(-24 * time.Hour)},
	{ID: 2, Title: "Drink tea", IsComplete: true, CreatedAt: time.Now().Add(-2 * time.Hour)},
	{ID: 3, Title: "Search job", IsComplete: false, CreatedAt: time.Now()},
}

func todoListHandler(w http.ResponseWriter, r *http.Request) {
	render := RenderCtx(r.Context())
	render.JSON(w, TodosResponse{&todoList})
}

func todoCreateHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	if err := todo.Validate(); err != nil {
		render.Status(http.StatusBadRequest).JSON(w, NewErrorValidate(err))
		return
	}
	todo.ID = len(todoList) + 1
	todo.CreatedAt = time.Now()
	todoList = append(todoList, todo)

	render.Status(http.StatusCreated).JSON(w, TodoResponse{&todo})
}

func todoGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	for _, todo := range todoList {
		if todo.ID == id {
			render.JSON(w, TodoResponse{&todo})
			return
		}
	}
	render.Status(http.StatusNotFound).JSON(w, NewError("Todo Not Found"))
}

func todoUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	for index, todo := range todoList {
		if todo.ID == id {
			json.NewDecoder(r.Body).Decode(&todo)
			if err := todo.Validate(); err != nil {
				render.Status(http.StatusBadRequest).JSON(w, NewErrorValidate(err))
				return
			}

			todoList[index] = todo

			render.JSON(w, TodoResponse{&todo})
			return
		}
	}
}

func todoDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	for index, todo := range todoList {
		if todo.ID == id {
			todoList = append(todoList[:index], todoList[index+1:]...)
			render.SendStatus(w, http.StatusOK)
			return
		}
	}
	render.Status(http.StatusNotFound).JSON(w, NewError("Todo Not Found"))
}
