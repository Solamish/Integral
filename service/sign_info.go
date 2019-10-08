package service

import "mobileSign/model"

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
	return data
}
