package controller

import (
	"github.com/gin-gonic/gin"
	"mobileSign/middleware"
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
	user.LastSignTime = time.Now().Local()
	user.Save()

	//c.Redirect(302,"https://wx.redrock.team/game?token="+token)
	c.JSON(200,gin.H{
		"msg": "ok",
	})
}


 

