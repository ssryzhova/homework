package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var task []Task

func main() {
	app := fiber.New()
	app.Get("/task", getTask)
	app.Get("/task/done", getDoneTask)
	app.Put("/task/:id/undone", unTask)
	app.Post("/task", createTask)
	app.Put("/task/:id", updateTask)
	app.Delete("/task/:id", deleteTask)
	app.Listen(":8080")
}
func getTask(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(task)
}

func getDoneTask(c *fiber.Ctx) error {
	var doneTask []Task
	for _, t := range task {
		if t.Done {
			doneTask = append(doneTask, t)
		}
	}
	return c.Status(fiber.StatusOK).JSON(doneTask)
}
func createTask(c *fiber.Ctx) error {
	var newTask Task
	if err := c.BodyParser(&newTask); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	newTask.ID = len(task) + 1
	newTask.Done = false
	task = append(task, newTask)
	return c.Status(fiber.StatusCreated).JSON(newTask)
}

func updateTask(c *fiber.Ctx) error {
	id := c.Params("id")

	for i := range task {
		if id == string(rune(task[i].ID)) {
			task[i].Done = true
			return c.Status(201).JSON(task)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
}
func unTask(c *fiber.Ctx) error {
	id := c.Params("id")
	for i := range task {
		if id == strconv.Itoa(task[i].ID) {
			task[i].Done = false
			return c.Status(fiber.StatusOK).JSON(task[i])
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
}
func deleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	for i := range task {
		if id == string(rune(task[i].ID)) {
			task = append(task[:i], task[i+1:]...)
			return c.Status(201).JSON(task)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
}
