package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"mobileSign/model"
	"mobileSign/service"
	"mobileSign/util/resps"
)

func SignInfo(c *gin.Context) {
	u, ok1 := c.Get("user")
	user, ok2 := u.(model.User)
	if !ok1 || !ok2 {
		log.Println("some error:", ok1, ok2)
		resps.DefinedError(c, resps.ParamError)
		return
	}
	data := service.SignInfo(user.RedId)
	resps.DefinedResp(c, resps.RespMsg{
		Status: 200,
		Info:   "success",
		Data:   data,
	})
}

 

