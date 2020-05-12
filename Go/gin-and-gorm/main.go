package main

import (
	//"net/http"

	"example.com/book/controllers"
	"example.com/book/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := models.SetupModels()

	// Provide db variable to controllers
	// we need to create a middleware that can provide the database instance to every single controller since they live in another file that can’t access the database instance directly.
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/books", controllers.FindBooks)        // new
	r.POST("/books", controllers.CreateBook)      // new
	r.GET("/books/:id", controllers.FindBook)     // new
	r.PATCH("/books/:id", controllers.UpdateBook) // new
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
