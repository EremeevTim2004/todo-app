package routes

import (
	"net/http"
	"todo-app/task-service/models"

	"github.com/gin-gonic/gin"
)

var tasks = make(map[string]models.Task)

func SetupRoutes(router *gin.Engine) {
	router.POST("/tasks", createTask)
	router.GET("/tasks/:id", getTask)
}

func createTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tasks[task.ID] = task
	c.JSON(http.StatusCreated, task)
}

func getTask(c *gin.Context) {
	id := c.Param("id")
	task, exists := tasks[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}
