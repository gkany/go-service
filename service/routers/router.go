package routers

import (
	"github.com/gin-gonic/gin"
	"go-service/service/controllers"
)

func InitRoute(r gin.IRouter) {
	hello(r.Group("/api/v1"))
}

func hello(api gin.IRouter) {
	helloGroup := api.Group("/hello")
	{
		helloGroup.POST("/world", controllers.World)
		helloGroup.GET("test", controllers.Test)
	}
}