package router

import (
	"github.com/brailyguzman/gotodo/internal/handlers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", handlers.HealthCheck)

	return r
}
