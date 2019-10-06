package resps

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RespMsg : http响应数据的通用结构
type RespMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
}

// NewRespMsg : 生成response对象
func NewRespMsg(code int, msg string, data interface{}) RespMsg {
	return RespMsg{
		Code: code,
		Msg:  msg,
	}
}


var (
	ParamError = RespMsg{
		Code: 10011,
		Msg:  "param error",

	}
	AuthorizedError = RespMsg{
		Code: 10011,
		Msg:  "Unauthorized",

	}
	Success = RespMsg{
		Code: 10000,
		Msg:  "success",

	}
)

//	自定义Error
func DefinedError(c *gin.Context, err RespMsg) {
	c.JSON(http.StatusOK, err)
}

func Unauthorized(c *gin.Context, err RespMsg) {
	c.JSON(http.StatusUnauthorized, err)
}

