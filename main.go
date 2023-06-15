package main

import (
	"belajar-golang-bookstore2/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Routing untuk buku
	books := r.Group("/books")
	{
		books.GET("/", handlers.GetBooks)
		books.GET("/:id", handlers.GetBook)
		books.POST("/", handlers.CreateBook)
		books.PUT("/:id", handlers.UpdateBook)
		books.DELETE("/:id", handlers.DeleteBook)
	}

	games := r.Group("/games")
	{
		games.GET("/headline", handlers.GetHeadlineGames)
		games.GET("/popular", handlers.GetPopularGames)
		games.GET("/trending", handlers.GetTrendingGames)
		games.GET("/details/:id", handlers.GetGameDetails)
		games.POST("/transactions", handlers.PostTransactions)
	}

	r.Run(":8000")
}
