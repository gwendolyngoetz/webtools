package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TextController struct{}

func (controller TextController) Get(context *gin.Context) {
	context.HTML(http.StatusOK, "text/index.tmpl", gin.H{
		"message": "get1",
	})
}

func (controller TextController) Get2(context *gin.Context) {
	context.HTML(http.StatusOK, "text/index.tmpl", gin.H{
		"message": "get2",
	})
}