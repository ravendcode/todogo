package todong

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Todo model
type Todo struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	IsComplete bool      `json:"isComplete"`
	CreatedAt  time.Time `json:"createdAt"`
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

var render = NewRender()

func todoListHandler(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, TodosResponse{&todoList})
}

func todoCreateHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		render.SetStatus(http.StatusBadRequest).JSON(w, NewError(err.Error()))
		return
	}
	todo.ID = len(todoList) + 1
	todo.CreatedAt = time.Now()
	todoList = append(todoList, todo)

	render.SetStatus(http.StatusCreated).JSON(w, TodoResponse{&todo})
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
	render.SetStatus(http.StatusNotFound).JSON(w, NewError("Todo Not Found"))
}

func todoUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	for index, todo := range todoList {
		if todo.ID == id {
			if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
				render.SetStatus(http.StatusBadRequest).JSON(w, NewError(err.Error()))
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
	render.SetStatus(http.StatusNotFound).JSON(w, NewError("Todo Not Found"))
}
