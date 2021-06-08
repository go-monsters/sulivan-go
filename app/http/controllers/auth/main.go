package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-monsters/sulivan/resources/views"
)

func GetLogin(c *gin.Context) {
	view := views.NewView("guest", "auth/login")
	view.Render(c.Writer, nil)
}

func Login(c *gin.Context) {

}

func GetRegister(c *gin.Context) {
	view := views.NewView("guest", "auth/register")
	view.Render(c.Writer, nil)
}

func Register(c *gin.Context) {
	view := views.NewView("guest", "auth/register")
	view.Render(c.Writer, nil)
}