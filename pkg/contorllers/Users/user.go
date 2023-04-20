package Users

import (
	"gin-jwt/pkg/config"
	"gin-jwt/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func SignupUser(user models.User) int64 {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println(err)
	}
	pwHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(pwHash)
	result := db.Table("Users").Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return result.RowsAffected
}

func LoginUser(data models.UserLogin) bool {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println(err)
	}
	var user models.User
	db.Table("Users").Where("email=?", data.Email).Find(&user)
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if passErr != nil {
		return false
	} else {
		return true
	}
}
