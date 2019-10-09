package api

import (
	"github.com/gin-gonic/gin"
	"mobileSign/model"
	"mobileSign/service"
	"mobileSign/util/resps"
)

type ItemForm struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Value    string `form:"value" json:"value" binding:"required"`
	Number   int    `form:"num" json:"num" binding:"required"`
	PhotoSrc string `form:"itemPhoto" json:"itemPhoto" binding:"required"`
}

func GetItemList(c *gin.Context) {
	itemList := service.GetItemList()
	resps.DefinedResp(c, resps.RespMsg{
		Status: 200,
		Info:   "success",
		Data:   itemList,
	})
}

func AddItem(c *gin.Context) {
	form := ItemForm{}
	if err := c.ShouldBind(&form); err != nil {
		resps.DefinedError(c, resps.ParamError)
		return
	}
	i := model.Item{
		Name:     form.Name,
		Value:    form.Value,
		Number:   form.Number,
		PhotoSrc: form.PhotoSrc,
	}
	if ok := i.AddItem(); !ok {
		resps.DefinedError(c, resps.ParamError)
		return
	}
	resps.DefinedResp(c, resps.RespMsg{
		Status: 200,
		Info:   "success",
		Data:   nil,
	})
}
