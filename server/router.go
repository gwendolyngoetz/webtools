package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.HTMLRender = CreateMultiRenderer("./templates")
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.StaticFS("/public", http.Dir("public"))

	AddRoutes(router)

	return router
}
