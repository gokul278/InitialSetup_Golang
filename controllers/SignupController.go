package controllers

import (
	service "AuthenticationService/Service"
	db "AuthenticationService/internal/DB"
	accesstoken "AuthenticationService/internal/Helper/AccessToken"
	model "AuthenticationService/internal/Model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PostSignupController() gin.HandlerFunc {
	return func(c *gin.Context) {

		var reqVal model.PostSignupNew

		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Something went wrong, Try Again",
			})

			return
		}

		dbConn, sqlDB := db.InitDB()
		defer sqlDB.Close()

		status := service.PostSignupService(dbConn, reqVal)

		if status {
			c.JSON(http.StatusOK, gin.H{
				"status":  true,
				"message": "Successfully Signed Up!",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  false,
				"message": "Something went wrong, Try Again",
			})
		}

	}
}

func GetSignupController() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, exists := c.Get("id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
			return
		}

		dbConn, sqlDB := db.InitDB()
		defer sqlDB.Close()

		getData := service.GetSignupService(dbConn)

		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"status":  true,
			"message": "Successfully Data Fetched",
			"data":    getData,
			"token":   accesstoken.CreateToken(id, 20*time.Minute),
		})
	}
}
