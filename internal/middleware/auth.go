package middleware

import (
	"github.com/brailyguzman/gotodo/db"
	"github.com/brailyguzman/gotodo/db/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User

		session := sessions.Default(ctx)
		userID := session.Get("user_id")

		if userID == nil {
			ctx.JSON(401, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		// Fetch the user from the database
		if err := db.DB.First(&user, userID).Error; err != nil {
			ctx.JSON(404, gin.H{"error": "User not found"})
			ctx.Abort()
			return
		}

		// Set the user in the context
		ctx.Set("user", user)
		ctx.Next()
	}
}
