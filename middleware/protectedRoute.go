package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
	"whatsgo/initializers"
	"whatsgo/models"
)

func RequireAuth(c *gin.Context) {
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//	Check expiration
		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			c.AbortWithStatus(401)
			return
		}

		//	Find user with token
		var user models.User
		if err := initializers.DB.Where("id = ?", claims["sub"]).First(&user).Error; err != nil {
			c.AbortWithStatus(401)
			return
		}

		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatus(401)
		return
	}

}
