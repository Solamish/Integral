package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"mobileSign/model"
	"mobileSign/service"
	"mobileSign/util/resps"
	"strconv"
)

func GetRecord(c *gin.Context) {
	pageStr := c.PostForm("page")
	sizeStr := c.PostForm("size")
	page, _ := strconv.Atoi(pageStr)
	size, _ := strconv.Atoi(sizeStr)
	u, ok1 := c.Get("user")
	user, ok2 := u.(model.User)
	if !ok1 || !ok2 {
		log.Println("some error:", ok1, ok2)
		resps.DefinedError(c, resps.ParamError)
		return
	}
	records := service.GetRecord(user.Stunum, page, size)
	resps.DefinedResp(c, resps.RespMsg{
		Status: 200,
		Info:   "success",
		Data:   records,
	})
}
