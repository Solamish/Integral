package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"mobileSign/model"
	"mobileSign/service"
	"mobileSign/util/resps"
	"net/http"
)

func GetBalance(c *gin.Context)  {
	u, ok1 := c.Get("user")
	user, ok2 := u.(model.User)
	if !ok1 || !ok2 {
		log.Println("some error:", ok1, ok2)
		resps.DefinedError(c, resps.ParamError)
		return
	}
	balance := service.Score(user.RedId)
	c.JSON(http.StatusOK, gin.H{
		"code": 10000,
		"balance": balance,
	})
}

 

