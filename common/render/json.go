package render

import (
	"github.com/gin-gonic/gin"
	"go-service/common/i18n"
	"go-service/common/status"
)

type RespJsonObj struct {
	Code      int         `json:"code"`
	Msg       string      `json:"message"`
	Reason    string      `json:"reason,omitempty""`
	Data      interface{} `json:"data"`
	RequestID string      `json:"request_id"`
}

func RespJson(c *gin.Context, code int, data interface{}) {
	result := &RespJsonObj{
		Code: code,
		Msg:  status.StatusText(code, i18n.GetLanguage(c)),
		Data: data,
	}
	c.JSON(status.HttpStatueCode(code), result)
}

func RespJsonWithError(c *gin.Context, code int, data interface{}, reason string) {
	result := &RespJsonObj{
		Code:   code,
		Msg:    status.StatusText(code, i18n.GetLanguage(c)),
		Reason: reason,
		Data:   data,
	}
	c.JSON(status.HttpStatueCode(code), result)
}

func RespJsonWithAbord(c *gin.Context, code int, data interface{}) {
	result := &RespJsonObj{
		Code: code,
		Msg:  status.StatusText(code, i18n.GetLanguage(c)),
		Data: data,
	}
	c.AbortWithStatusJSON(status.HttpStatueCode(code), result)
}