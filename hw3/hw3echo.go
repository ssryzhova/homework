package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var task []Task

func main() {
	e := echo.New()
	e.GET("/task", getTask)
	e.GET("/task/done", getDoneTask)
	e.PUT("/task/:id/undone", unTask)
	e.POST("/task", createTask)
	e.PUT("/task/:id", updateTask)
	e.DELETE("/task/:id", deleteTask)
	e.Logger.Fatal(e.Start(":8080"))
}
func getTask(c echo.Context) error {
	return c.JSON(http.StatusOK, task)
}

func getDoneTask(c echo.Context) error {
	var doneTask []Task
	for _, t := range task {
		if t.Done {
			doneTask = append(doneTask, t)
		}
	}
	return c.JSON(http.StatusOK, doneTask)
}
func createTask(c echo.Context) error {
	var newTask Task
	if err := c.Bind(&newTask); err != nil {
		return err
	}
	newTask.ID = len(task) + 1
	newTask.Done = false
	task = append(task, newTask)
	return c.JSON(http.StatusCreated, newTask)
}

func updateTask(c echo.Context) error {
	id := c.Param("id")

	for i := range task {
		if id == strconv.Itoa(task[i].ID) {
			task[i].Done = false
			return c.JSON(http.StatusOK, task[i])
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
}
func unTask(c echo.Context) error {
	id := c.Param("id")
	for i := range task {
		if id == strconv.Itoa(task[i].ID) {
			task[i].Done = false
			return c.JSON(http.StatusOK, task[i])
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")

	for i := range task {
		if id == strconv.Itoa(task[i].ID) {
			task = append(task[:i], task[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
}
