package main

import (
	"todo-app/auth-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8081") // Запуск на порту 8081
}
