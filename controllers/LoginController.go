package controllers

import (
	service "AuthenticationService/Service"
	db "AuthenticationService/internal/DB"
	model "AuthenticationService/internal/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostLoginController() gin.HandlerFunc {

	return func(c *gin.Context) {

		var reqVal model.LoginReq

		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Something went wrong, Try Again",
			})
			return
		}

		dbConn, sqlDB := db.InitDB()
		defer sqlDB.Close()

		resVal := service.PostLoginService(dbConn, reqVal)

		response := gin.H{
			"status":  resVal.Status,
			"message": resVal.Message,
		}

		if resVal.Token != "" {
			response["token"] = resVal.Token
		}

		c.JSON(http.StatusOK, response)
	}
}
