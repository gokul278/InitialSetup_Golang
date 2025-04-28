package routes

import (
	"AuthenticationService/controllers"

	"github.com/gin-gonic/gin"
)

func InitLoginRoutes(router *gin.Engine) {
	route := router.Group("/v1/login")
	route.POST("/", controllers.LoginController())
}
