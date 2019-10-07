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

	QAGroup := router.Group("/QA")
	{
		IntergralGroup := QAGroup.Group("/Intergral")
		{
			IntergralGroup.POST("/checkIn", controller.Sign)
		}

		UserGroup := QAGroup.Group("/User")
		{
			UserGroup.GET("/getScoreStatus",controller.GetBalance)
		}

	}
	return router
}
