package bootstrap

import (
	"github.com/gin-gonic/gin"
	"sullivan/config"
	"sullivan/routes"
)

var Router *gin.Engine

func Start() {
	//load config from env file
	config.LoadConfig()
	//create new gin
	Router = gin.Default()
	//register views
	Router.LoadHTMLGlob("resources/views/*.tmpl")
	//Router.LoadHTMLGlob("resources/views/*/*.tmpl")
	//register routes
	//if production go run in production
	//if debug -> turn on debug
	routes.Register(Router)
	//logger
	// 		log on file or any other platforms. log daily
	//middleware
		// exception handler
		//response handler
		// request response logger
		//use uuid for uniq request
		//formatter for error and json
	//auth
	//crypt with key in env
	//cors
	//CSRF
	//db
	//		use GORM and postgres. use any database and multi connection. user select default connection. support multi connection
	//load template
	//add webpack vue bootstrap and ...
}
