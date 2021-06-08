package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-monsters/sulivan/app/http/controllers"
	"github.com/go-monsters/sulivan/app/http/controllers/auth"
)

func WebRoute(router *gin.Engine) {
	router.GET("/", controllers.Welcome)
	router.GET("/login", auth.Login)
	router.GET("/register", auth.Register)
}
