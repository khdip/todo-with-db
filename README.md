# todo-with-db

This repo contains a CRUD todo web application where you can create, view, edit and delete tasks.

For database I used PostgreSQL. The connection was made using a go library named [sqlx](https://github.com/jmoiron/sqlx).

To Create a database run the below in psql shell:
```CREATE DATABASE todo;```

If you have a different database please modify the below code in the main.go file.
```db, err := sqlx.Connect("postgres", "user=postgres password=Anubis0912 dbname=todo sslmode=disable")```
