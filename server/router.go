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

	homeController := new(controllers.HomeController)

	router.GET("/", homeController.Index)

	return router
}
