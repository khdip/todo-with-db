package handler

import (
	"net/http"
)

type FormData struct {
	Todo   ToDo
	Errors map[string]string
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	ErrorValue := map[string]string{}
	todo := ToDo{}
	h.LoadCreatedForm(w, todo, ErrorValue)
}

func (h *Handler) StoreTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	myTask := r.FormValue("Task")
	todo := ToDo{
		Task: myTask,
	}
	if myTask == "" {
		ErrorValue := map[string]string{
			"Task": "This field can not be empty.",
		}
		h.LoadCreatedForm(w, todo, ErrorValue)
		return
	} else if len(myTask) < 3 {
		ErrorValue := map[string]string{
			"Task": "This field should have atleast 3 characters",
		}
		h.LoadCreatedForm(w, todo, ErrorValue)
		return
	}
	const insertTodo = `INSERT INTO tasks(title, is_completed) VALUES($1, $2);`
	res := h.db.MustExec(insertTodo, myTask, false)
	ok, err := res.RowsAffected()
	if err != nil || ok == 0 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/todo", http.StatusTemporaryRedirect)
}

func (h *Handler) CompleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/todo/complete/"):]
	if id == "" {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}
	const completeTodo = `UPDATE tasks SET is_completed = true WHERE id=$1`
	res := h.db.MustExec(completeTodo, id)
	ok, err := res.RowsAffected()
	if err != nil || ok == 0 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/todo", http.StatusTemporaryRedirect)
}

func (h *Handler) EditTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/todo/edit/"):]
	if id == "" {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	const getTodo = `SELECT * FROM tasks WHERE id=$1`
	var todo ToDo
	h.db.Get(&todo, getTodo, id)

	if todo.ID == 0 {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}
	h.LoadEditForm(w, todo, map[string]string{})
}

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/todo/update/"):]
	if id == "" {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	const getTodo = `SELECT * FROM tasks WHERE id=$1`
	var todo ToDo
	h.db.Get(&todo, getTodo, id)

	if todo.ID == 0 {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	myNewTask := r.FormValue("Task")
	todo.Task = myNewTask
	if myNewTask == "" {
		ErrorValue := map[string]string{
			"Task": "This field can not be empty.",
		}
		h.LoadEditForm(w, todo, ErrorValue)
		return
	} else if len(myNewTask) < 3 {
		ErrorValue := map[string]string{
			"Task": "This field should have atleast 3 characters",
		}
		h.LoadEditForm(w, todo, ErrorValue)
		return
	}

	const updateTodo = `UPDATE tasks SET title = $2 WHERE id=$1`
	res := h.db.MustExec(updateTodo, id, myNewTask)
	ok, err := res.RowsAffected()
	if err != nil || ok == 0 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/todo", http.StatusTemporaryRedirect)
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/todo/delete/"):]
	if id == "" {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	const getTodo = `SELECT * FROM tasks WHERE id=$1`
	var todo ToDo
	h.db.Get(&todo, getTodo, id)

	if todo.ID == 0 {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}
	const deleteTodo = `DELETE FROM tasks WHERE id=$1`
	res := h.db.MustExec(deleteTodo, id)
	ok, err := res.RowsAffected()
	if err != nil || ok == 0 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/todo", http.StatusTemporaryRedirect)
}

// Form Validation
func (h *Handler) LoadCreatedForm(w http.ResponseWriter, todo ToDo, myErrors map[string]string) {
	form := FormData{
		Todo:   todo,
		Errors: myErrors,
	}

	err := h.templates.ExecuteTemplate(w, "create-todo.html", form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) LoadEditForm(w http.ResponseWriter, todo ToDo, myErrors map[string]string) {
	form := FormData{
		Todo:   todo,
		Errors: myErrors,
	}

	err := h.templates.ExecuteTemplate(w, "edit-todo.html", form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
