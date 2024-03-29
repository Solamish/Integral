package middleware

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"mobileSign/model"
	"mobileSign/util/resps"
	"mobileSign/util/signs"
	"strings"
)

type UserInfo struct {
	RedID      string `json:"redid"`
	HeadImgUrl string `json:"headImgUrl"`
	Nickname   string `json:"nickname"`
	RealName   string `json:"realName"`
	StuNum     string `json:"stuNum"`
}

// idnum和stunum鉴权
func OldAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		stuNum := c.PostForm("stunum")
		idNum := c.PostForm("idnum")
		if stuNum == "" || idNum == "" {
			resps.DefinedError(c, resps.AuthorizedError)
			c.Abort()
			return
		}
		if !signs.Verify(stuNum, idNum) {
			resps.DefinedError(c, resps.AuthorizedError)
			c.Abort()
			return
		}
		c.Set("stuNum", stuNum)
		c.Next()
	}
}

// magicloop鉴权
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//从header取出token
		token := c.Request.Header.Get("token")
		method := c.Request.Method

		if method == "OPTIONS" {
			c.Next()
		}

		user, err := CheckToken(token)
		if err != nil {
			resps.Unauthorized(c, resps.AuthorizedError)
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

func CheckToken(token string) (user model.User, err error) {
	token = strings.ReplaceAll(token, " ", "+")
	tokenSlice := strings.Split(token, ".")
	if len(tokenSlice) != 2 {
		err = errors.New("token error")
		return
	}
	payload := tokenSlice[0]
	//signature := tokenSlice[1]
	//if ok := signs.CheckToken(payload, signature); !ok {
	//	err = errors.New("token error")
	//	return
	//}
	b, _ := base64.StdEncoding.DecodeString(payload)
	u := UserInfo{}
	_ = json.Unmarshal(b, &u)

	//将头像url的协议由http改为https
	headImgUrlSlice := strings.Split(u.HeadImgUrl, ":")
	head := headImgUrlSlice[0] + "s:"
	u.HeadImgUrl = head + headImgUrlSlice[1]

	user = model.User{
		RedId:      u.RedID,
		NickName:   u.Nickname,
		HeadImgUrl: u.HeadImgUrl,
		Stunum:     u.StuNum,
		UserName:   u.RealName,
	}

	return
}
