package server

import (
	"gwendolyngoetz/webtool/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.StaticFS("/public", http.Dir("public"))
	router.LoadHTMLGlob("templates/**/*")

	addRoutes(router)

	return router
}

func addRoutes(router *gin.Engine) {
	addHomeControllerRoutes(router)
	addTextControllerRoutes(router)
}

func addHomeControllerRoutes(router *gin.Engine) {
	controller := new(controllers.HomeController)
	router.GET("/", controller.Index)
}

func addTextControllerRoutes(router *gin.Engine) {
	controller := new(controllers.TextController)

	group := router.Group("text")
	{
		group.GET("/get1", controller.Get)
		group.GET("/get2", controller.Get2)
	}
}
