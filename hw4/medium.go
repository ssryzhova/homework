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

type Password string
type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var task []Task
var user []User
var nextTaskId = 1
var nextUserId = 1

func main() {
	r := gin.Default()
	r.GET("/task", getTask)
	r.GET("/user", getUser)
	r.GET("/task/:id", getTaskid)
	r.GET("/user/:id", getUserid)
	r.POST("/task", createTask)
	r.POST("/user", createUser)
	r.PUT("/task/:id", updateTask)
	r.PUT("/user/:id", updateUser)
	r.DELETE("/task/:id", deleteTask)
	r.DELETE("/user/:id", deleteUser)
	r.Run(":8080")
}

// task
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
	newTask.ID = nextTaskId
	nextTaskId++
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

// user
func getUser(c *gin.Context) {
	c.JSON(http.StatusOK, user)
}

func getUserid(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, u := range user {
		if u.ID == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func createUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newUser.ID = nextUserId
	nextUserId++

	user = append(user, newUser)

	c.JSON(http.StatusCreated, newUser)
}

func updateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var updated User
	if err := c.BindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	for i, u := range user {
		if u.ID == id {
			updated.ID = id
			user[i] = updated
			c.JSON(http.StatusOK, updated)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func deleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, u := range user {
		if u.ID == id {
			user = append(user[:i], user[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
