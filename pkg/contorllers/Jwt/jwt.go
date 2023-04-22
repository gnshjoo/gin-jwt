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

func FetchToken(token utils.Token) (bool, error) {
	accessToken := token.AccessToken

	db, err := config.ConnectDatabase()
	if err != nil {
		return false, err
	}
	var savedToken utils.Token
	db.Table("Tokens").Where("access_token = ?", accessToken).Scan(&savedToken)

	_, err = utils.ValidToken(savedToken.AccessToken)
	if err != nil {
		return false, err
	}
	_, err = utils.ValidToken(savedToken.RefreshToken)
	if err != nil {
		return false, err
	}

	return true, nil

}
