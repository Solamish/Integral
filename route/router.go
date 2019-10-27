package route

import (
	"github.com/gin-gonic/gin"
	"mobileSign/api"
	"mobileSign/middleware"
)

func Load(router *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	// 入口
	router.GET("/user/enter", api.Enter)

	router.Use(middleware.AuthMiddleware())
	router.Use(middleware.CorsMiddleware())

	// QA
	QAGroup := router.Group("/QA")
	{
		IntergralGroup := QAGroup.Group("/Intergral")
		{
			IntergralGroup.POST("/checkIn", api.Sign)
			IntergralGroup.POST("/getItemList", api.GetItemList)
			IntergralGroup.POST("/addItem", api.AddItem)
		}

		UserGroup := QAGroup.Group("/User")
		{
			UserGroup.POST("/getScoreStatus", api.SignInfo)
			UserGroup.POST("/integralRecords", api.GetRecord)
		}

	}

	// home
	HomeGroup := router.Group("/Home")
	{
		PersonGroup := HomeGroup.Group("/Person")
		{
			PersonGroup.POST("/search", api.UserInfo)
		}
	}
	return router
}
