package resps

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


type ErrorMsg struct {
	Status int         `json:"status"`
	Info  string      `json:"info"`
}

// RespMsg : http响应数据的通用结构
type RespMsg struct {
	Status int         `json:"status"`
	Info  string      `json:"info"`
	Data  interface{} `json:"data"`
}



var (
	ParamError = ErrorMsg{
		Status: 10011,
		Info:  "param error",

	}
	AuthorizedError = ErrorMsg{
		Status: 10011,
		Info:  "Unauthorized",

	}
	Resp = RespMsg{
		Status: 200,
		Info:   "success",
		Data:   nil,
	}
)

//	自定义Error
func DefinedError(c *gin.Context, err ErrorMsg) {
	c.JSON(http.StatusOK, err)
}

func DefinedResp(c *gin.Context, resp RespMsg)  {
	c.JSON(http.StatusOK, resp)
}

func Unauthorized(c *gin.Context, err ErrorMsg) {
	c.JSON(http.StatusUnauthorized, err)
}

