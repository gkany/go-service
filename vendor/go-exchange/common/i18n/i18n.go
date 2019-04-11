package i18n

import (
	"github.com/gin-gonic/gin"
	ln "golang.org/x/text/language"
)


const (
	CN = 0
	EN = 1
)

func GetLanguage(c *gin.Context) int {
	//Accept-Language: zh-CN,zh;q=0.9,en;q=0.8

	accept_language := c.Request.Header.Get("Accept-Language")
	tags, _, _ := ln.ParseAcceptLanguage(accept_language)
	if len(tags) > 0 {
		if tags[0].String() == "zh-CN" || tags[0].String() == "zh" {
			return CN
		}
	}
	return CN
}


