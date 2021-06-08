package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-monsters/sulivan/resources/views"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Welcome(c *gin.Context) {
	view := views.NewView("welcome", "welcome")
	view.Render(c.Writer, nil)
}
