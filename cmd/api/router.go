package api

import (
	"book-category-api/internal/handlers"
	"book-category-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/users/login", handlers.UserLogin)

		categories := api.Group("/categories")
		categories.Use(middleware.JWTAuthMiddleware())
		{
			categories.GET("", handlers.GetCategories)
			categories.POST("", handlers.CreateCategory)
			categories.GET("/:id", handlers.GetCategory)
			categories.DELETE("/:id", handlers.DeleteCategory)
			categories.GET("/:id/books", handlers.GetBooksByCategory)
		}

		books := api.Group("/books")
		books.Use(middleware.JWTAuthMiddleware())
		{
			books.GET("", handlers.GetBooks)
			books.POST("", handlers.CreateBook)
			books.GET("/:id", handlers.GetBook)
			books.DELETE("/:id", handlers.DeleteBook)
		}
	}

	return r
}
