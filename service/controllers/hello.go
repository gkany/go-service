package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/liyue201/go-logger"
	"go-service/common/render"
	"go-service/common/status"
)

func World(c *gin.Context) {
	logger.Debugf("[World] http server: echo hello world")

	render.RespJson(c, status.OK, "hello world from http server with http post")
	return
}

func Test(c *gin.Context) {
	logger.Debugf("[World] http server: echo hello world")

	render.RespJson(c, status.OK, "hello world from http server with http get")
	return
}