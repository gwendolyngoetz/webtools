package controllers

import (
	"html"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TextController struct{}

func (controller TextController) GetHtmlEncode(context *gin.Context) {
	context.HTML(http.StatusOK, "text_htmlencode", gin.H{})
}

func (controller TextController) PostHtmlEncode(context *gin.Context) {
	textarea1, _ := context.GetPostForm("textarea1")

	cleantextarea1 := html.EscapeString(textarea1)

	context.HTML(http.StatusOK, "text_htmlencodepartial", gin.H{
		"text": cleantextarea1,
	})
}

func (controller TextController) PostHtmlDecode(context *gin.Context) {
	textarea1, _ := context.GetPostForm("textarea1")

	cleantextarea1 := html.UnescapeString(textarea1)

	context.HTML(http.StatusOK, "text_htmlencodepartial", gin.H{
		"text": cleantextarea1,
	})
}
