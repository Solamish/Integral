package route

import (
	"github.com/gin-gonic/gin"
	"mobileSign/controller"
	"mobileSign/middleware"
)

func Load(router *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	// 入口
	router.GET("/user/enter", controller.Enter)

	router.Use(middleware.AuthMiddleware())
	router.Use(middleware.CorsMiddleware())

	QAgroup := router.Group("/QA")
	{
		QAgroup.POST("/Integral/checkIn", controller.Sign)
	}
	return router
}
