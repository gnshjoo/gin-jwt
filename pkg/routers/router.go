package routers

import (
	"encoding/json"
	"fmt"
	"gin-jwt/pkg/contorllers/Users"
	"gin-jwt/pkg/models"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

// RequestPing
// @name ping
// @Summary This is simple Summary
// @Description This is detail Description
// @Accept json
// @Produce json
// @Router /ping [get]
// @Success 200 {object} string
func RequestPing(c *gin.Context) {
	fmt.Println("got ping")
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}

// SignUp
// @name create user
// @summary
// @Param name path string true "name"
// @Param password path string true "password"
// @Param email path string true "email"
// @Description
// @Accept json
// @Produce json
// @Router /signup [post]
// @Success 200 {object} string
func SignUp(c *gin.Context) {
	body := c.Request.Body
	value, err := io.ReadAll(body)
	if err != nil {
		log.Println(err)
	}
	var data models.User
	err = json.Unmarshal([]byte(value), &data)
	if err != nil {
		log.Println(err)
	}

	result := Users.SignupUser(data)
	c.JSON(200, gin.H{"message": result})
}

// Login
// @name login user
// @summary
// @Param password path string true "password"
// @Param email path string true "email"
// @Description
// @Accept json
// @Produce json
// @Router /login [post]
// @Success 200 {object} string
func Login(c *gin.Context) {
	body := c.Request.Body
	var data models.UserLogin
	err := json.NewDecoder(body).Decode(&data)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"message": err})
	}

	result := Users.LoginUser(data)
	c.JSON(200, gin.H{"message": result})
}
