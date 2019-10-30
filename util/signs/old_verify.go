package signs

import (
	"encoding/json"
	"mobileSign/util/https"
)

var (
	verifyUrl = "https://wx.redrock.team/api/verify"
)

type VerifyResp struct {
	Info   string
	Status int
}

func Verify(stuNum, idNum string) bool {
	var params map[string]string
	params = make(map[string]string, 2)
	params["stuNum"] = stuNum
	params["idNum"] = idNum
	body := https.SendPost(verifyUrl, params)
	var verifyResp VerifyResp
	_ = json.Unmarshal(body, &verifyResp)
	if verifyResp.Status == -1 {
		return false
	}
	return true
}
 

