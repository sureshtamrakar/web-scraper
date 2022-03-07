package routes

import (
	"github.com/gin-gonic/gin"
	controller_fetch "github.com/sureshtamrakar/web-scraper/controller/fetch"
	controller_login "github.com/sureshtamrakar/web-scraper/controller/login"
	controller_register "github.com/sureshtamrakar/web-scraper/controller/register"
)

func AddRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/login", controller_login.Login)            // login user
	r.POST("/register", controller_register.CreateUser) // register user
	r.POST("/fetch", controller_fetch.Create)           // fetch url
	r.GET("/fetch/:url", controller_fetch.Get)          // get value from stored url
	r.GET("/fetch/date", controller_fetch.Date)         // get value from start to end
	r.GET("/fetch", controller_fetch.List)              // list of scraped value
	r.GET("/fetch/page/:id", controller_fetch.Paginate) // paginate

	return r
}
