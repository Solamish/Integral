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
			IntergralGroup.POST("/checkIn", api.Sign)   // 签到
			IntergralGroup.POST("/getItemList", api.GetItemList)  // 获取商品
			IntergralGroup.POST("/addItem", api.AddItem)	// 添加商品
		}

		UserGroup := QAGroup.Group("/User")
		{
			UserGroup.POST("/getScoreStatus", api.SignInfo)   // 签到信息
			UserGroup.POST("/integralRecords", api.GetRecord)	// 积分使用记录
		}

	}

	// home
	HomeGroup := router.Group("/Home")
	{
		PersonGroup := HomeGroup.Group("/Person")
		{
			PersonGroup.POST("/search", api.UserInfo)   // 获取个人信息
		}
	}
	return router
}
