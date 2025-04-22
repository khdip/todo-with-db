package main

import (
	"fmt"
	"log"
	"net/http"

	"practice/todo-with-db/handler"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	var schema = `
			CREATE TABLE IF NOT EXISTS tasks (
				id serial,
				title text,
				is_completed boolean,

				primary key(id)
			);`

	db, err := sqlx.Connect("postgres", "user=postgres password=Anubis0912 dbname=todo sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(schema)

	h := handler.GetHandler(db)
	http.HandleFunc("/todo", h.GetTodo)
	http.HandleFunc("/todo/create", h.CreateTodo)
	http.HandleFunc("/todo/store", h.StoreTodo)
	http.HandleFunc("/todo/complete/", h.CompleteTodo)
	http.HandleFunc("/todo/edit/", h.EditTodo)
	http.HandleFunc("/todo/Update/", h.UpdateTodo)
	http.HandleFunc("/todo/delete/", h.DeleteTodo)
	fmt.Println("Server Starting...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server Not Found", err)
	}
}
