package main

import (
	"log"
	"os"

	"github.com/brailyguzman/gotodo/db"
	"github.com/brailyguzman/gotodo/internal/handlers"
	"github.com/brailyguzman/gotodo/internal/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))

	r.Use(sessions.Sessions("gotodo_session", store))

	api := r.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)

		auth := api.Group("/auth")
		{
			auth.POST("/signup", handlers.CreateUser)
			auth.POST("/login", handlers.LoginUser)
			auth.GET("/me", handlers.GetCurrentUser)
			auth.POST("/logout", middleware.AuthMiddleware(), handlers.LogoutUser)
		}

		todos := api.Group("/todos", middleware.AuthMiddleware())
		{
			todos.POST("/", handlers.CreateTodo)
			todos.GET("/", handlers.GetTodos)
			todos.PATCH("/:id", handlers.EditTodo)
			todos.DELETE("/:id", handlers.DeleteTodo)
		}

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
