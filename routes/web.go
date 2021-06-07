package routes

import (
	"github.com/gin-gonic/gin"
	"sullivan/app/http/controllers"
)

func WebRoute(router *gin.Engine) {
	router.GET("/", controllers.Welcome)
}
