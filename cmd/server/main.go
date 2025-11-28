package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nvrbckdown/bookstore-api-gateway/config"
	"github.com/nvrbckdown/bookstore-api-gateway/handlers"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Set Gin mode
	gin.SetMode(cfg.GinMode)

	// Initialize router
	router := gin.New()

	// Apply middleware
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())

	// Register routes
	registerRoutes(router, cfg)

	// Start server
	log.Printf("Server starting on port %s\n", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func registerRoutes(router *gin.Engine, cfg *config.Config) {
	// Health check
	router.GET("/health", handlers.HealthCheck)

	// API version group
	v1 := router.Group("/api/v1")
	{
		// Book service routes
		books := v1.Group("/books")
		{
			books.GET("", handlers.GetBooks(cfg.BookstoreServiceURL))
			books.GET("/:id", handlers.GetBook(cfg.BookstoreServiceURL))
			books.POST("", handlers.CreateBook(cfg.BookstoreServiceURL))
			books.PUT("/:id", handlers.UpdateBook(cfg.BookstoreServiceURL))
			books.DELETE("/:id", handlers.DeleteBook(cfg.BookstoreServiceURL))
		}

		// Author routes
		authors := v1.Group("/authors")
		{
			authors.GET("", handlers.GetAuthors(cfg.BookstoreServiceURL))
			authors.GET("/:id", handlers.GetAuthor(cfg.BookstoreServiceURL))
			authors.POST("", handlers.CreateAuthor(cfg.BookstoreServiceURL))
			authors.PUT("/:id", handlers.UpdateAuthor(cfg.BookstoreServiceURL))
			authors.DELETE("/:id", handlers.DeleteAuthor(cfg.BookstoreServiceURL))
			authors.GET("/:id/books", handlers.GetAuthorBooks(cfg.BookstoreServiceURL))
		}

		// Order service routes
		orders := v1.Group("/orders")
		{
			orders.GET("", handlers.GetOrders(cfg.OrderServiceURL))
			orders.GET("/:id", handlers.GetOrder(cfg.OrderServiceURL))
			orders.POST("", handlers.CreateOrder(cfg.OrderServiceURL))
			orders.PUT("/:id", handlers.UpdateOrder(cfg.OrderServiceURL))
			orders.DELETE("/:id", handlers.DeleteOrder(cfg.OrderServiceURL))
			orders.GET("/customer/:email", handlers.GetCustomerOrders(cfg.OrderServiceURL))
		}
	}
}
