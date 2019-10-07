package service

import "mobileSign/model"

// 获取积分余额
func Score(redId string) int {
	var balance int
	balance =  model.GetScore(redId)
	return balance
}

 

