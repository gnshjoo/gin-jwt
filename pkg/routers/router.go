package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
