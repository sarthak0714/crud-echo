package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sarthak0714/crud-echo/model"
)

func GetTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.GetTask(db))
	}
}

func AddTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var t model.Task
		t.Completed = false
		c.Bind(&t)
		op, err := model.AddTask(db, t.Desc, t.Completed)
		if err == nil {
			return c.JSON(http.StatusCreated, op)
		} else {
			return err
		}
	}
}

func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		op, err := model.DeleteTask(db, id)
		if err == nil {
			return c.JSON(http.StatusOK, op)
		} else {
			return err
		}
	}
}

func UpdateTaskStatus(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))	
		op, err := model.UpdateStatus(db, id)
		if err == nil {
			return c.JSON(http.StatusCreated, op)
		} else {
			return err
		}
	}
}
