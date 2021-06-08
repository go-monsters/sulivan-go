package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-monsters/sulivan/app/http/controllers"
	"github.com/go-monsters/sulivan/app/http/controllers/auth"
)

func WebRoute(router *gin.Engine) {
	router.GET("/", controllers.Welcome)
	router.GET("/login", auth.GetLogin)
	router.GET("/register", auth.GetRegister)
	router.POST("/login", auth.Login)
	router.POST("/register", auth.Register)
}

