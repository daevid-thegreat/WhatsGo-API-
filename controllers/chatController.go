package controllers

import (
	"github.com/gin-gonic/gin"
)

func OpenChat(c *gin.Context) {
	// Get the user email from the request body
	var body struct {
		Email string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(400, gin.H{"error": "Fields are empty or not valid"})
		return
	}

	c.JSON(200, gin.H{"message": "User created successfully"})

}
