package Users

import (
	"gin-jwt/pkg/config"
	"gin-jwt/pkg/models"
	"log"
)

func SignupUser(user models.User) int64 {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println(err)
	}

	result := db.Table("Users").Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return result.RowsAffected
}

func LoginUser(data models.UserLogin) models.User {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println(err)
	}
	var result models.User
	db.Table("Users").Where("email=?", data.Email).Scan(&result)

	println(&result)
	return result
}
