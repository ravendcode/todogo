package todong

import (
	"time"
)

var todoList = Todos{
	{ID: 1, Title: "Create chat", IsComplete: false, CreatedAt: time.Now().Add(-24 * time.Hour)},
	{ID: 2, Title: "Drink tea", IsComplete: true, CreatedAt: time.Now().Add(-2 * time.Hour)},
	{ID: 3, Title: "Search job", IsComplete: false, CreatedAt: time.Now()},
}

func todoCreateRepo(todo *Todo) {
	todo.ID = len(todoList) + 1
	todo.CreatedAt = time.Now()
	todoList = append(todoList, *todo)
}

func todoFindRepo(id int) (int, Todo) {
	for index, todo := range todoList {
		if todo.ID == id {
			return index, todo
		}
	}
	return -1, Todo{}
}

func todoDestroyRepo(index int) {
	todoList = append(todoList[:index], todoList[index+1:]...)
}
