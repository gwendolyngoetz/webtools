package controllers

import (
	"gwendolyngoetz/webtool/forms"
	"html"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TextController struct{}

func (controller TextController) GetHtmlEncode(context *gin.Context) {
	context.HTML(http.StatusOK, "text_htmlencode", gin.H{})
}

func (controller TextController) PostHtmlEncode(context *gin.Context) {
	var formdata forms.HtmlEncodeForm
	context.Bind(&formdata)

	convertedText := ""

	switch aaa := formdata.Action; aaa {
	case "encode":
		convertedText = html.EscapeString(formdata.TextArea1)
	case "decode":
		convertedText = html.UnescapeString(formdata.TextArea1)
	}

	context.HTML(http.StatusOK, "text_htmlencodepartial", gin.H{
		"text": convertedText,
	})
}

func (controller TextController) PostHtmlDecode(context *gin.Context) {
	var formdata forms.HtmlEncodeForm
	context.Bind(&formdata)
	cleantextarea1 := html.UnescapeString(formdata.TextArea1)

	context.HTML(http.StatusOK, "text_htmlencodepartial", gin.H{
		"text": cleantextarea1,
	})
}
