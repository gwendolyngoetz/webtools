package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TextController struct{}

func (controller TextController) Get(context *gin.Context) {
	context.HTML(http.StatusOK, "text_index", gin.H{
		"message": "get1",
	})
}

func (controller TextController) Get2(context *gin.Context) {
	context.HTML(http.StatusOK, "text_index", gin.H{
		"message": "get2",
	})
}

func (controller TextController) GetHtmlEncode(context *gin.Context) {
	context.HTML(http.StatusOK, "text_htmlencode", gin.H{})
}

func (controller TextController) GetHtmlDecode(context *gin.Context) {
	context.HTML(http.StatusOK, "text_htmldecode", gin.H{})
}
