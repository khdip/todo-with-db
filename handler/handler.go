package handler

import (
	"html/template"

	"github.com/jmoiron/sqlx"
)

type ToDo struct {
	ID          int    `db:"id" json:"id"`
	Task        string `db:"title" json:"task"`
	IsCompleted bool   `db:"is_completed" json:"is_completed"`
}

type Handler struct {
	templates *template.Template
	db        *sqlx.DB
}

func GetHandler(db *sqlx.DB) *Handler {
	hand := &Handler{
		db: db,
	}
	hand.GetTemplate()
	return hand
}

func (h *Handler) GetTemplate() {
	h.templates = template.Must(template.ParseFiles("templates/create-todo.html", "templates/list-todo.html", "templates/edit-todo.html"))
}
