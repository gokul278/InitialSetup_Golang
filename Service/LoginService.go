package service

import (
	accesstoken "AuthenticationService/internal/Helper/AccessToken"
	hashdb "AuthenticationService/internal/Helper/HashDB"
	model "AuthenticationService/internal/Model"
	"fmt"
	"os"
	"time"

	"gorm.io/gorm"
)

func PostLoginService(db *gorm.DB, reqVal model.LoginReq) model.LoginResponse {

	dbkey := os.Getenv("ENCRYPT_DB")

	var user []model.Userdata

	result := db.Where("username = ?", reqVal.Username).First(&user)

	if result.Error != nil {
		fmt.Println("Get Error:", result.Error)
		return model.LoginResponse{
			Status:  false,
			Message: "Something went wrong, Try Again",
		}
	}

	if len(user) > 0 {

		if hashdb.Decrypt(user[0].Password, dbkey) == reqVal.Password {

			return model.LoginResponse{
				Status:  true,
				Message: "Successfully Logined In",
				Token:   accesstoken.CreateToken(int(user[0].ID), 20*time.Minute),
			}
		} else {
			return model.LoginResponse{
				Status:  false,
				Message: "Invalid Login Credential",
			}
		}

	} else {
		return model.LoginResponse{
			Status:  false,
			Message: "Invalid Login Credential",
		}

	}

}
