package service

import (
	accesstoken "AuthenticationService/internal/Helper/AccessToken"
	becrypt "AuthenticationService/internal/Helper/Becrypt"
	hashdb "AuthenticationService/internal/Helper/HashDB"
	logger "AuthenticationService/internal/Helper/Logger"
	model "AuthenticationService/internal/Model"
	"AuthenticationService/query"
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
				// Token:   accesstoken.CreateToken(int(user[0].ID), 20*time.Minute),
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

func LoginServices(db *gorm.DB, reqVal model.LoginReq) model.LoginResponse {
	log := logger.InitLogger()
	var AdminLoginModel []model.AdminLoginModel

	// Execute the raw SQL query with the username (phone number)
	err := db.Raw(query.LoginAdminSQL, reqVal.Username).Scan(&AdminLoginModel).Error
	if err != nil {
		log.Error("LoginService DB Error: " + err.Error())
		return model.LoginResponse{
			Status:  false,
			Message: "Internal server error",
		}
	}

	// Check if any user found
	if len(AdminLoginModel) == 0 {
		log.Warn("LoginService Invalid Credentials(u) for Username: " + reqVal.Username)
		return model.LoginResponse{
			Status:  false,
			Message: "Invalid username or password",
		}
	}

	// Password verification
	user := AdminLoginModel[0]
	match := becrypt.ComparePasswords(user.ADHashPass, reqVal.Password)

	if !match {
		log.Warn("LoginService Invalid Credentials(p) for Username: " + reqVal.Username)
		return model.LoginResponse{
			Status:  false,
			Message: "Invalid username or password",
		}
	}

	// If matched
	log.Info("LoginService Logined Successfully for Username: " + reqVal.Username)

	history := model.RefTransHistory{
		TransTypeId: 1,
		THData:      hashdb.Encrypt("Logged In Successfully"),
		UserId:      user.UserId,
		THActionBy:  user.UserId,
	}

	errhistory := db.Create(&history).Error
	if errhistory != nil {
		log.Error("LoginService INSERT ERROR at Trnasaction: " + errhistory.Error())
		return model.LoginResponse{
			Status:  false,
			Message: "Internal server error",
		}
	}
	return model.LoginResponse{
		Status:  true,
		Message: "Login successful",
		Token:   accesstoken.CreateToken(user.UserId, user.RId, "0", 20*time.Minute),
	}
}
