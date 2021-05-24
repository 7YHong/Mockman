package routes

import (
	"github.com/gin-gonic/gin"
	"mockman/configor"
	"mockman/controllers"
	"mockman/middleware"
)

func SetRoutes(router *gin.Engine) {
	router.Use(
		gin.Recovery(), // error handle
		middleware.RespLogMiddleware(),
	)

	router.GET("hello",
		middleware.CorsMiddleware(),
		func(ctx *gin.Context) {
			hello := controllers.HelloController{}
			hello.Index(ctx)
		},
	)

	router.Any("echo", func(c *gin.Context) {
		reqInfo, ok := c.Get("reqInfo")
		if !ok {
			return
		}
		c.JSON(200, reqInfo)
	})

	api:=router.Group("api")
	for _,mock:=range configor.Config.Mock{
		api.Handle(mock.Method,mock.Path,newApiHandle(mock.Response))
	}

}

func newApiHandle(resp string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Data(200,"application/json; charset=utf-8", []byte(resp))
	}
}