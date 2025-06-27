package handlers

import (
	"strconv"

	"github.com/brailyguzman/gotodo/db"
	"github.com/brailyguzman/gotodo/db/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type CreateTodoRequest struct {
	Text string `json:"text"`
}

type EditTodoRequest struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func CreateTodo(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	uid, ok := userID.(uint)

	if !ok {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var todo CreateTodoRequest

	if err := ctx.BindJSON(&todo); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	newTodo := models.Todo{
		Text:   todo.Text,
		UserID: uid,
	}

	if err := newTodo.Create(db.DB, &newTodo); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create todo"})
		return
	}

	ctx.JSON(201, gin.H{"message": "Todo created successfully", "id": newTodo.ID})
}

func EditTodo(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	uid, ok := userID.(uint)

	if !ok {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var input EditTodoRequest

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var todo models.Todo

	if err := db.DB.First(&todo, uint(id)).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	if todo.UserID != uid {
		ctx.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	todo.Text = input.Text
	todo.Done = input.Done

	if err := db.DB.Save(&todo).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Updated successfully"})
}

func DeleteTodo(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("user_id")
	uid, ok := userID.(uint)

	if !ok {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var todo models.Todo

	if err := db.DB.First(&todo, uint(id)).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	if todo.UserID != uid {
		ctx.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	if err := todo.Delete(db.DB, &todo); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete todo"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Deleted Successfully"})
}
