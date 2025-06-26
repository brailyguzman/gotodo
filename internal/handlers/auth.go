package handlers

import (
	"github.com/brailyguzman/gotodo/db"
	"github.com/brailyguzman/gotodo/db/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(ctx *gin.Context) {
	var user CreateUserRequest

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var existingUser models.User

	if err := db.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		ctx.JSON(400, gin.H{"error": "User already exists"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	newUser := models.User{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: string(hash),
	}

	if err := newUser.Create(db.DB, &newUser); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	// TODO: Encrypt user ID in cookie
	session := sessions.Default(ctx)
	session.Set("user_id", newUser.ID)
	if err := session.Save(); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to save session"})
		return
	}

	ctx.JSON(201, gin.H{"message": "User created successfully", "user": UserResponse{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
	}})
}

func GetCurrentUser(ctx *gin.Context) {
	userID, err := ctx.Cookie("user_id")

	if err != nil {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User

	if err := db.DB.First(&user, userID).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(200, UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}

func LoginUser(ctx *gin.Context) {
	var user LoginUserRequest
	var err error

	if err = ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var foundUser models.User

	if err = db.DB.Where("email = ?", user.Email).First(&foundUser).Error; err != nil {
		ctx.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(user.Password)); err != nil {
		ctx.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	session := sessions.Default(ctx)
	session.Set("user_id", foundUser.ID)

	if err = session.Save(); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to save session"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Login successful", "user": UserResponse{
		ID:    foundUser.ID,
		Name:  foundUser.Name,
		Email: foundUser.Email,
	}})
}

func LogoutUser(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()

	if err := session.Save(); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to clear session"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Logout successful"})
}
