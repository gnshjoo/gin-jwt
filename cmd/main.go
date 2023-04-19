package main

import (
	_ "gin-jwt/docs"
	"gin-jwt/pkg/routers"
	_ "gin-jwt/pkg/routers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func setupSwagger(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "swagger/index.html")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func main() {
	r := gin.Default()
	setupSwagger(r)
	r.GET("/ping", routers.RequestPing)
	r.POST("/signup", routers.SignUp)

	err := r.Run()
	if err != nil {
		return
	}
}
