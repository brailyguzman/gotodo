package main

import (
	"log"
	"os"

	"github.com/brailyguzman/gotodo/db"
	"github.com/brailyguzman/gotodo/internal/handlers"
	"github.com/brailyguzman/gotodo/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", handlers.HealthCheck)
	auth := r.Group("/auth")
	{
		auth.POST("/signup", handlers.CreateUser)
		auth.POST("/login", handlers.LoginUser)
		auth.POST("/logout", handlers.LogoutUser)
		auth.GET("/me", handlers.GetCurrentUser)
	}

	// use middleware with group for todos
	todos := r.Group("/todos", middleware.AuthMiddleware())
	{
		// TODO: Implement the handlers for todos
	}

	return r
}

func main() {
	var err error

	err = godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("POSTGRES_DSN")

	if dsn == "" {
		log.Fatal("POSTGRES_DSN environment variable is not set")
	}

	db.ConnectDatabase(dsn)
	db.Migrate()

	r := NewRouter()
	err = r.Run(":3000")

	if err != nil {
		log.Fatalf("Error running the server on PORT 3000: %v", err)
	}
}
