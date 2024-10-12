package main

import (
	"todo-app/task-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8082") // Запуск на порту 8082
}
