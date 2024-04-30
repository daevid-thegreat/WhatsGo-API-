package controllers

import (
	"github.com/gin-gonic/gin"
	"whatsgo/initializers"
	"whatsgo/models"
)

func OpenChat(c *gin.Context) {
	// Get the other user email from the request body
	var body struct {
		Email string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(400, gin.H{"error": "Fields are empty or not valid"})
		return
	}

	var user models.User
	if err := initializers.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "No account found"})
		return
	}

	// Get the logged-in user
	var loggedInUser models.User
	if err := initializers.DB.Where("email = ?", c.GetString("email")).First(&loggedInUser).Error; err != nil {
		c.JSON(400, gin.H{"error": "No account logged in"})
		return

	}

	// Create chat

	c.JSON(200, gin.H{"message": "User created successfully"})

}
