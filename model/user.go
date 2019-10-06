package model

import (
	"log"
	"time"
)

type User struct {
	ID           int       `gorm:"primary_key"`
	RedId        string    `gorm:"column:redid"`
	NickName     string    `gorm:"column:nickname"`
	HeadImgUrl   string    `gorm:"column:head_img_url"`
	LastSignTime time.Time `gorm:"column:last_sign_time"`
	ContDays     int       `gorm:"column:cont_days"`
	Score        int       `gorm:"column:score"`
}

func (user *User) Save() {
	tmp := User{}
	DB.Where("redid = ?", user.RedId).Find(&tmp)
	if tmp.RedId == "" {
		DB.Create(user)
	}
}

func GetLastSignTime(redId string) (lastSignTime time.Time) {
	err := DB.Raw("select last_sign_time from users where redid = ?", redId).Row().Scan(&lastSignTime)
	if err != nil {
		log.Println("fail to get last sign time", err)
	}
	return
}

func GetContDays(redId string) (contDays int) {
	err := DB.Raw("select cont_days from users where redid = ?", redId).Row().Scan(&contDays)
	if err != nil {
		log.Println("fail to get contDays", err)
	}
	return
}

func ResetContDays(redId string, now time.Time) {
	//err := DB.Table("users").Where("redid = ?", redId).Updates(map[string]interface{}{"cont_days": 1, "last_sign_time": now,}).Error
	err := DB.Exec("UPDATE users SET cont_days = 1, last_sign_time = ?  WHERE redid = ?", now, redId).Error
	if err != nil {
		log.Println("fail to reset cont days", err)
	}



}

func UpdateContDays(redId string, now time.Time) {
	err := DB.Exec("UPDATE users SET cont_days = cont_days + 1, last_sign_time = ? WHERE redid = ?", now, redId).Error
	if err != nil {
		log.Println("fail to update cont days", err)
	}
}

func UpdateScore(redId string, score int) {
	err := DB.Exec("UPDATE users SET score = score + ? WHERE redid = ?", score, redId).Error
	if err != nil {
		log.Println("fail to update score", err)
	}
}
