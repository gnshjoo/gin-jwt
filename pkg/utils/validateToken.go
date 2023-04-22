package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type AccessDetails struct {
	AccessUuid string
	ID         string
}

func ValidateToken(token string) (*jwt.Token, error) {
	verifyToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error")
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	return verifyToken, nil
}

func ValidToken(tokenString string) (*AccessDetails, error) {
	token, err := ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		id := fmt.Sprintf("%v", claims["user_id"])
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			accessUuid,
			id,
		}, nil
	}
	return nil, err
}
