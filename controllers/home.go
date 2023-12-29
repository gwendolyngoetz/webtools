package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (controller HomeController) Index(context *gin.Context) {
	context.HTML(http.StatusOK, "home/index.tmpl", gin.H{
		"otherthing": "hi",
	})
}
