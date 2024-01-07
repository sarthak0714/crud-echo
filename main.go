package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sarthak0714/crud-echo/handler"
)

func main() {

	db := initDatabase("tasks.db")
	migrate(db)

	e := echo.New()
	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Working\n /tasks {POST}\n /tasks/:id {PUT,DELETE}")
	})
	e.GET("/tasks",handler.GetTask(db))
	e.POST("/tasks",handler.AddTask(db))
	e.PUT("/tasks/:id",handler.UpdateTaskStatus(db))
	e.DELETE("/tasks/:id",handler.DeleteTask(db))

	e.Logger.Fatal(e.Start(":42069"))
}

func initDatabase(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if db == nil {
		panic("DB not found")
	}
	if err != nil {
		panic(err)
	}

	return db
}

func migrate(db *sql.DB) {
	qry := `
	create table if not exists tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		desc VARCHAR NOT NULL,
		completed BOOLEAN NOT NULL
	)
	`
	_, err := db.Exec(qry)

	if err != nil {
		panic(err)
	}
}
