package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var task []Task

func main() {
	r := gin.Default()
	r.GET("/task", getTask)
	r.GET("/task/done", getDoneTask)
	r.PUT("/task/:id/undone", unTask)
	r.POST("/task", createTask)
	r.PUT("/task/:id", updateTask)
	r.DELETE("/task/:id", deleteTask)
	r.Run(":8080")
}
func getTask(c *gin.Context) {
	c.JSON(http.StatusOK, task)
}

func getDoneTask(c *gin.Context) {
	var doneTask []Task
	for _, t := range task {
		if t.Done {
			doneTask = append(doneTask, t)
		}
	}
	c.JSON(http.StatusOK, doneTask)
}
func createTask(c *gin.Context) {
	var newTask Task
	if err := c.BindJSON(&newTask); err != nil {
		return
	}
	newTask.ID = len(task) + 1
	newTask.Done = false
	task = append(task, newTask)
	c.JSON(http.StatusCreated, newTask)
}

func updateTask(c *gin.Context) {
	id := c.Param("id")

	for i := range task {
		if id == string(rune(task[i].ID)) {
			task[i].Done = true
			c.JSON(http.StatusOK, task[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
func unTask(c *gin.Context) {
	id := c.Param("id")
	for i := range task {
		if id == strconv.Itoa(task[i].ID) {
			task[i].Done = false
			c.JSON(http.StatusOK, task[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
func deleteTask(c *gin.Context) {
	id := c.Param("id")

	for i := range task {
		if id == string(rune(task[i].ID)) {
			task = append(task[:i], task[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
}
