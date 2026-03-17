package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Status string

const (
	New        Status = "Новая"
	InProgress Status = "В процессе"
	Done       Status = "Завершена"
)

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      `json:"status"`
}

var task []Task
var nextId = 1

func main() {
	r := gin.Default()
	r.GET("/task", getTask)
	r.GET("/task/:id", getTaskid)
	r.POST("/task", createTask)
	r.PUT("/task/:id", updateTask)
	r.DELETE("/task/:id", deleteTask)
	r.Run(":8080")
}
func getTask(c *gin.Context) {
	c.JSON(http.StatusOK, task)
}

func getTaskid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, t := range task {
		if t.ID == id {
			c.JSON(http.StatusOK, t)
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	}
}
func createTask(c *gin.Context) {
	var newTask Task
	if err := c.BindJSON(&newTask); err != nil {
		return
	}
	newTask.ID = nextId
	nextId++
	if newTask.Status == "" {
		newTask.Status = New
	}

	task = append(task, newTask)
	c.JSON(http.StatusCreated, newTask)
}

func updateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var updated Task
	if err := c.BindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	for i, t := range task {
		if t.ID == id {
			updated.ID = id
			task[i] = updated
			c.JSON(http.StatusOK, updated)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func deleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, t := range task {
		if t.ID == id {
			task = append(task[:i], task[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
