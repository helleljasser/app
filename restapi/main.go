// main.go
package main

import (
	"restapi/api"
	"restapi/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	repository.InitDB()

	// Create a new Gin router
	router := gin.Default()

	// Routes
	router.GET("/users", api.GetUsers)
	router.GET("/users/:id", api.GetUser)
	router.POST("/users", api.CreateUser)
	router.PUT("/users/:id", api.UpdateUser)
	router.DELETE("/users/:id", api.DeleteUser)

	// Run the application on port 8080
	router.Run(":8080")
}
