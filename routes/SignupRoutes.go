package routes

import (
	"AuthenticationService/controllers"
	accesstoken "AuthenticationService/internal/Helper/AccessToken"

	"github.com/gin-gonic/gin"
)

func InitSignupRoutes(router *gin.Engine) {
	route := router.Group("/v1/signup")
	route.POST("/new", controllers.PostSignupController())
	route.GET("/", accesstoken.JWTMiddleware(), controllers.GetSignupController())
}
