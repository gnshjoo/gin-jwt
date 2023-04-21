package Users

import (
	"gin-jwt/pkg/config"
	"gin-jwt/pkg/contorllers/Jwt"
	"gin-jwt/pkg/models"
	"gin-jwt/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func SignupUser(user models.User) (int64, error) {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	pwHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(pwHash)
	result := db.Table("Users").Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func LoginUser(data models.UserLogin) (*utils.Token, error) {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println(err)
	}
	var user models.User
	db.Table("Users").Where("email=?", data.Email).Find(&user)
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if passErr != nil {
		return nil, passErr
	} else {
		token, err := utils.GetJwtToken(user.ID)
		if err != nil {
			return nil, err
		}
		token.ID = data.Email
		err = Jwt.SaveToken(token)
		if err != nil {
			return nil, err
		}
		return token, nil
	}
}
