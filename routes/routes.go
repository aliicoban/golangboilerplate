package routes

import (
	"github.com/alicobanserver/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.POST("/signUp", controllers.Signup)
	router.POST("/signIn", controllers.Signin)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To Test API",
	})
	return
}
