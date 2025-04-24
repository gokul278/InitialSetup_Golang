package service

import (
	hashdb "AuthenticationService/internal/Helper/HashDB"
	model "AuthenticationService/internal/Model"
	"fmt"
	"os"

	"gorm.io/gorm"
)

func PostSignupService(db *gorm.DB, reqVal model.PostSignupNew) bool {

	dbkey := os.Getenv("ENCRYPT_DB")

	user := model.Userdata{
		Username:  reqVal.Email,
		Password:  hashdb.Encrypt(reqVal.Password, dbkey),
		CreatedBy: "Admin",
	}

	dbResult := db.Create(&user)

	if dbResult.Error != nil {
		fmt.Println("Insertion Failed", dbResult.Error)
		return false
	}

	return true
}

func GetSignupService(db *gorm.DB) []model.UserdataRes {

	// dbkey := os.Getenv("ENCRYPT_DB")

	var users []model.UserdataRes

	result := db.Raw(`SELECT "id","username","created_at","created_by","updated_at","updated_by" FROM public."userdata"`).Scan(&users)
	// result := db.Find(&users)
	if result.Error != nil {
		fmt.Println("Get Error:", result.Error)
		return nil
	}

	// for i := range users {
	// 	users[i].Username = hashdb.Decrypt(users[i].Username, dbkey)
	// 	users[i].UpdatedBy = hashdb.Decrypt(users[i].UpdatedBy, dbkey)
	// 	users[i].CreatedBy = hashdb.Decrypt(users[i].CreatedBy, dbkey)
	// }

	return users
}
