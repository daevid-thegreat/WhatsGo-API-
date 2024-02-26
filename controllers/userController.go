package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
	"whatsgo/initializers"
	"whatsgo/models"
)

func SignUp(c *gin.Context) {
	// Get the user email and password from the request body
	var body struct {
		Email    string
		Password string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(400, gin.H{"error": "Fields are empty or not valid"})
		return
	}

	// Check if the user already exists in the database

	if err := initializers.DB.Where("email = ?", body.Email).First(&models.User{}).Error; err == nil {
		c.JSON(400, gin.H{"error": "Email already exists"})
		return
	}

	// Hash password

	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error hashing password"})
		return
	}

	// Create user

	user := models.User{
		Email:    body.Email,
		Password: string(password),
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(200, gin.H{"message": "User created successfully"})

}

func SignIn(c *gin.Context) {
	// Get the user email and password from the request body
	var body struct {
		Email    string
		Password string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(400, gin.H{"error": "Fields are empty or not valid"})
		return
	}

	// Check if the user exists in the database

	var user models.User
	if err := initializers.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "Email or password is incorrect"})
		return
	}

	// Check if the password is correct

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(400, gin.H{"error": "Email or password is incorrect"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(500, gin.H{"error": "Error signing token"})
		return
	}

	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "/", "localhost", false, true)
	c.JSON(200, gin.H{"message": "User signed in successfully"})

}

func ForgotPassword(c *gin.Context) {
	// Get the user email from the request body
	var body struct {
		Email string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(400, gin.H{"error": "Fields are empty or not valid"})
		return
	}

	// Check if the user exists in the database

	var user models.User
	if err := initializers.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "Email is incorrect"})
		return
	}
}

func ResetPassword(c *gin.Context) {
	// Get the user email and password from the request body
	var body struct {
		Email       string
		OTP         string
		NewPassword string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(400, gin.H{"error": "Fields are empty or not valid"})
		return
	}

	// Check if the user exists in the database

	var user models.User
	if err := initializers.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "Email is incorrect"})
		return
	}

	// Hash password

	password, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error hashing password"})
		return
	}

	// Update user password

	if err := initializers.DB.Model(&user).Update("password", string(password)).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error updating password"})
		return
	}

	c.JSON(200, gin.H{"message": "Password reset successfully"})

}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(200, gin.H{"user": user})
}
