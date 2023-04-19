package Users

import (
	"gin-jwt/pkg/config"
	"gin-jwt/pkg/models"
	"gorm.io/gorm"
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

func LoginUser(email, password string) *gorm.DB {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println(err)
	}

	result := db.Table("Users").Where("email=?", email)

	if result.Error != nil {
		log.Println(result.Error)
	}
	println(result)
	return result
}
