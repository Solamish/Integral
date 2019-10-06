package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"mobileSign/model"
	"mobileSign/service"
	"mobileSign/util/resps"
	"net/http"
)

func Sign(c *gin.Context) {
	//token := c.Request.Header.Get("token")
	u, ok1 := c.Get("user")
	user, ok2 := u.(model.User)
	if !ok1 || !ok2 {
		log.Println("some error:", ok1, ok2)
		resps.DefinedError(c, resps.ParamError)
		return
	}
	score, signTime, isSigned := service.Sign(user.RedId)

	c.JSON(http.StatusOK, gin.H{
		"code":10000,
		"interal": score,
		"lastSignTime": signTime.Format("2006-01-02 15:04:05"),
		"isSigned": isSigned,
	})
}
 

