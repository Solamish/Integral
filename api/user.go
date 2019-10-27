package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"mobileSign/middleware"
	"mobileSign/model"
	"mobileSign/service"
	"mobileSign/util/resps"
	"time"
)

func Enter(c *gin.Context) {
	token := c.Query("token")
	user, err := middleware.CheckToken(token)

	if err != nil || &user.RedId == nil{
		resps.DefinedError(c, resps.ErrorMsg{
			Status: 12345,
			Info:  err.Error(),
		})
		return
	}
	// TODO 要改呀
	user.LastSignTime = time.Now().Local()
	user.Save()

	//c.Redirect(302,"https://wx.redrock.team/game?token="+token)
	c.JSON(200,gin.H{
		"msg": "ok",
	})
}

func UserInfo(c *gin.Context) {
	u, ok1 := c.Get("user")
	user, ok2 := u.(model.User)
	if !ok1 || !ok2 {
		log.Println("some error:", ok1, ok2)
		resps.DefinedError(c, resps.ParamError)
		return
	}
	userInfo := service.UserInfo(user.RedId)
	resps.DefinedResp(c, resps.RespMsg{
		Status: 200,
		Info:   "success",
		Data:   userInfo,
	})
}


 

