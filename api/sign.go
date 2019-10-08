package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"mobileSign/model"
	"mobileSign/service"
	"mobileSign/util/resps"
	"net/http"
)

type Data struct {
	data interface{}
}

func Sign(c *gin.Context) {
	//token := c.Request.Header.Get("token")
	u, ok1 := c.Get("user")
	user, ok2 := u.(model.User)
	if !ok1 || !ok2 {
		log.Println("some error:", ok1, ok2)
		resps.DefinedError(c, resps.ParamError)
		return
	}
	score := service.Sign(user.RedId)
	if score == -1 {
		c.JSON(http.StatusOK, gin.H{
			"info":   "today had checked in",
			"status": 403,
		})
		return
	} else if score == -2 {
		c.JSON(http.StatusOK, gin.H{
			"info":   "Not within the specified date",
			"status": 405,
		})
		return
	}
	resps.DefinedResp(c, resps.Resp)
}
