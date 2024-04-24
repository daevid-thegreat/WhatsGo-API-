package main

import (
	"github.com/gin-gonic/gin"
	"whatsgo/controllers"
	"whatsgo/initializers"
	"whatsgo/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDb()
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.ForwardedByClientIP = true
	proxyErr := r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.2", "10.0.0.0/8"})
	if proxyErr != nil {
		return
	}
	// USer Endpoints
	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.SignIn)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/forgot-password", controllers.ForgotPassword)
	r.POST("/reset-password", controllers.ResetPassword)
	r.POST("/update-user", middleware.RequireAuth, controllers.UpdateUser)
	// Chat Endpoint
	err := r.Run()
	if err != nil {
		return
	}
}
