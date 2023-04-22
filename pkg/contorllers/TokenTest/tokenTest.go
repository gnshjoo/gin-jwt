package TokenTest

import (
	"gin-jwt/pkg/contorllers/Jwt"
	"gin-jwt/pkg/utils"
)

func TokenVerifyTest(token utils.Token) error {
	fetchToken, err := Jwt.FetchToken(token)
	if err != nil {
		return err
	}
	println(fetchToken)
	return nil
}
