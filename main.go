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
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.ForwardedByClientIP = true
	proxyErr := r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.2", "10.0.0.0/8", "whatsgo.onrender.com"})
	if proxyErr != nil {
		return
	}
	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.SignIn)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	err := r.Run()
	if err != nil {
		return
	}
}
