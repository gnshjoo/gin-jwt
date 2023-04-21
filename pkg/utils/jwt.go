package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
	"time"
)

type Token struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

var jwtKey = []byte("JWT Key 값")

func GetJwtToken(userId int64) (*Token, error) {
	tk := &Token{}
	tk.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	tk.AccessUuid = uuid.NewV4().String()

	tk.RtExpires = time.Now().Add(time.Hour * 24).Unix()
	tk.RefreshUuid = uuid.NewV4().String()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["access_uuid"] = tk.AccessUuid
	atClaims["exp"] = tk.AtExpires

	var err error

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tk.AccessToken, err = at.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = tk.RefreshUuid
	rtClaims["user_id"] = userId
	rtClaims["exp"] = tk.RtExpires

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	tk.RefreshToken, err = rt.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	return tk, nil

}
