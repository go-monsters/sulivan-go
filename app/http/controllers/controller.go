package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}


func Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.tmpl", gin.H{
		"title": "Main website",
	})
}

