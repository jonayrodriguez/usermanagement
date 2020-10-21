package controller

import (
	"net/http"

	"github.com/jonayrodriguez/usermanagement/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/jonayrodriguez/usermanagement/pkg/model"
)

//TODO- Implement Pagination (keep in mind, that gorm is creating the transaction for us)

// FindUsers - Find all users
func FindUsers(c *gin.Context) {
	var users []model.User
	database.GetInstance().Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUser - Get a user by username
func GetUser(c *gin.Context) {
	// Get model if exist
	var user model.User
	result := database.GetInstance().Where("username = ?", c.Param("username")).First(&user)
	// We could use c.JSON(http.StatusNoContent, nil), but it always adds something to the response
	if err := result.Error; err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// CreateUser - Create a user
func CreateUser(c *gin.Context) {
	// Validate input
	var input model.CreateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := model.User{Username: input.Username, Surname: input.Surname, Email: input.Email}
	database.GetInstance().Create(&user)

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// DeleteUser - Delete a user
func DeleteUser(c *gin.Context) {
	// Get model if exist
	var user model.User
	result := database.GetInstance().Where("username = ?", c.Param("username")).First(&user)
	// We could use c.JSON(http.StatusNoContent, nil), but it always adds something to the response
	if err := result.Error; err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	database.GetInstance().Delete(&user)
	c.Writer.WriteHeader(http.StatusNoContent)

}
