package handlers

import (
	"github.com/brailyguzman/gotodo/db"
	"github.com/brailyguzman/gotodo/db/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type CreateTodoRequest struct {
	Text string `json:"text"`
}

func CreateTodo(ctx *gin.Context) {
	var todo CreateTodoRequest

	if err := ctx.BindJSON(&todo); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Assuming the user ID is stored in the session
	session := sessions.Default(ctx)
	userID := session.Get("user_id")

	if userID == nil {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	newTodo := models.Todo{
		Text:   todo.Text,
		UserID: userID.(uint),
	}

	if err := newTodo.Create(db.DB, &newTodo); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create todo"})
		return
	}

	ctx.JSON(201, gin.H{"message": "Todo created successfully"})
}

// TODO
func EditTodo(ctx *gin.Context) {
}
