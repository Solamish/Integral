package service

import (
	"mobileSign/model"
	"mobileSign/util/times"
)

type Data struct {
	Integral     int `json:"integral"`
	CheckInDays  int `json:"check_in_days"`
	IsCheckToday int `json:"is_check_today"`
}


func SignInfo(redId string) (Data){
	// TODO is_check_today
	data := Data{}
	data.Integral = model.GetIntegral(redId)
	data.CheckInDays  = model.GetContDays(redId)

	lastSignTime := model.GetLastSignTime(redId)
	now,_ ,_ ,_ := times.GetTimeInfo()
	lastSignTimeStr := lastSignTime.Format("2006-01-02")
	nowStr := now.Format("2006-01-02")

	if lastSignTimeStr == nowStr {
		data.IsCheckToday = 1
	}
	return data
}
