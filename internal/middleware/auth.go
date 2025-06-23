package middleware

import (
	"github.com/brailyguzman/gotodo/db"
	"github.com/brailyguzman/gotodo/db/models"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		// Get the user ID from the context
		userID, err := c.Cookie("user_id")

		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Fetch the user from the database
		if err := db.DB.First(&user, userID).Error; err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Set the user in the context
		c.Set("user", user)
		c.Next()
	}
}
