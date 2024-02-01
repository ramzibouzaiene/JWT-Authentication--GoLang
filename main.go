package main

import (
	"auth-jwt/controllers"
	"auth-jwt/initializers"
	"auth-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadVariables()
	initializers.ConnectDb()
	initializers.SyncDatabase()
}

func main() {

	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
