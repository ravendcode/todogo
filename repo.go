package main

// var todoList = Todos{
// 	{ID: 1, Title: "Create chat", IsComplete: false, CreatedAt: time.Now().Add(-24 * time.Hour)},
// 	{ID: 2, Title: "Drink tea", IsComplete: true, CreatedAt: time.Now().Add(-2 * time.Hour)},
// 	{ID: 3, Title: "Search job", IsComplete: false, CreatedAt: time.Now()},
// }

// TodoRepo struct
type TodoRepo struct {
}

// List todo
func (t TodoRepo) List() *Todos {
	todos := &Todos{}
	db.Order("id asc").Find(todos)
	return todos
}

// Create todo
func (t TodoRepo) Create(todo *Todo) {
	db.Create(todo)
}

// Find todo
func (t TodoRepo) Find(todo *Todo, id int) {
	db.First(todo, id)
}

// Save todo
func (t TodoRepo) Save(todo *Todo) {
	db.Save(todo)
}

// Destroy todo
func (t TodoRepo) Destroy(todo *Todo) {
	db.Delete(todo)
}

// // Find todo
// func (t TodoRepo) Find(id int) (int, *Todo) {
// 	for index, todo := range todoList {
// 		if todo.ID == int64(id) {
// 			return index, &todo
// 		}
// 	}
// 	return -1, &Todo{}
// }

// // Create todo
// func (t TodoRepo) Create(todo *Todo) {
// 	if len(todoList) == 0 {
// 		todo.ID = 1
// 	} else {
// 		lastTodo := todoList[len(todoList)-1]
// 		todo.ID = lastTodo.ID + 1
// 	}
// 	todo.CreatedAt = time.Now()
// 	todoList = append(todoList, *todo)
// }

// // Destroy todo
// func (t TodoRepo) Destroy(index int) {
// 	todoList = append(todoList[:index], todoList[index+1:]...)
// }
