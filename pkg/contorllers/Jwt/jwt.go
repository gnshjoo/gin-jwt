package Jwt

import (
	"gin-jwt/pkg/config"
	"gin-jwt/pkg/utils"
	"log"
)

func SaveToken(token *utils.Token) error {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println(err)
		return err
	}
	tk := db.Table("Tokens").Create(&token)
	if tk.Error != nil {
		return tk.Error
	}

	return nil
}

func GetToken(rtoken, atoken string) (utils.Token, error) {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println(err)
		return utils.Token{}, err
	}

	var token utils.Token
	tk := db.Table("Tokens").Where("access_token=?", atoken).Where("refresh_token=?", rtoken).Find(&token)
	if tk.Error != nil {
		return utils.Token{}, err
	}
	return token, nil
}
