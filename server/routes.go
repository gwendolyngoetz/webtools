package server

import (
	"gwendolyngoetz/webtool/controllers"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine) {
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
		group.GET("/html-encode", controller.GetHtmlEncode)
		group.GET("/html-decode", controller.GetHtmlDecode)
		group.POST("/html-encode1", controller.PostHtmlEncode)
	}
}
