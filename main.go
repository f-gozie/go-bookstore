package main

import (
	"github.com/gin-gonic/gin"

	"rest-go/models"
	"rest-go/controllers"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.FetchBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run("localhost:8001")
}