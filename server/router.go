package server

import (
	"gwendolyngoetz/webtool/controllers"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.HTMLRender = createMultiRenderer("./templates")
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.StaticFS("/public", http.Dir("public"))

	addRoutes(router)

	return router
}

func createMultiRenderer(templatesDir string) multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/**/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	log.Printf(" -----------------------------")
	log.Printf("| TEMPLATE NAMES")
	log.Printf("|-----------------------------")

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)

		templateName := strings.Replace(include, "templates/includes/", "", 1)
		log.Printf("| %s", templateName)

		renderer.AddFromFiles(templateName, files...)
	}

	log.Printf(" ------------------------------")

	return renderer
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
