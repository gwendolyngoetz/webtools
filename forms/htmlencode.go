package forms

type HtmlEncodeForm struct {
	TextArea1 string `form:"textarea1"`
	Action    string `form:"action"`
}
