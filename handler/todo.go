package handler

import (
	"net/http"
)

type ListToDo struct {
	ToDo_list []ToDo
}

func (h *Handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	todos := []ToDo{}
	h.db.Select(&todos, "SELECT * FROM tasks")
	lt := ListToDo{ToDo_list: todos}
	err := h.templates.ExecuteTemplate(w, "list-todo.html", lt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
