package api

import (
	"github.com/gin-gonic/gin"
	"mobileSign/service"
	"mobileSign/util/resps"
)

func GetItemList(c *gin.Context) {
	itemList := service.GetItemList()
	resps.DefinedResp(c,resps.RespMsg{
		Status: 200,
		Info:   "success",
		Data:   itemList,
	})
}

 

