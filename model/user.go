package model

import (
	"log"
	"time"
)

type User struct {
	ID                int       `gorm:"primary_key"`
	RedId             string    `gorm:"column:redid"`
	NickName          string    `gorm:"column:nickname"`
	HeadImgUrl        string    `gorm:"column:head_img_url"`
	LastSignTime      time.Time `gorm:"column:last_sign_time"`
	CheckInDays       int       `gorm:"column:check_in_days"`
	Integral          int       `gorm:"column:integral"`
	Phone             string    `gorm:"column:phone"`
	Gendor            string    `gorm:"column:gendor"`
	Stunum            string    `gorm:"column:stunum"`
	Introduction      string    `gorm:"column:introduction"`
	QQ                string    `gorm:"column:qq"`
	UserName          string    `gorm:"column:username"`
	PhotoThumbnailSrc string    `gorm:"column:photo_thumbnail_src"`   // 缩略图
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
	err := DB.Raw("select check_in_days from users where redid = ?", redId).Row().Scan(&contDays)
	if err != nil {
		log.Println("fail to get check_in_days", err)
	}
	return
}

func ResetContDays(redId string, now time.Time) {
	//err := DB.Table("users").Where("redid = ?", redId).Updates(map[string]interface{}{"cont_days": 1, "last_sign_time": now,}).Error
	err := DB.Exec("UPDATE users SET check_in_days = 1, last_sign_time = ?  WHERE redid = ?", now, redId).Error
	if err != nil {
		log.Println("fail to reset check_in_days days", err)
	}

}

func UpdateContDays(redId string, now time.Time) {
	err := DB.Exec("UPDATE users SET check_in_days = check_in_days + 1, last_sign_time = ? WHERE redid = ?", now, redId).Error
	if err != nil {
		log.Println("fail to update check_in_days days", err)
	}
}

func UpdateIntegral(redId string, integral int) {
	err := DB.Exec("UPDATE users SET integral = integral + ? WHERE redid = ?", integral, redId).Error
	if err != nil {
		log.Println("fail to update integral", err)
	}
}

func GetIntegral(redId string) (totalIntegral int) {
	err := DB.Raw("select integral from users where redid = ?", redId).Row().Scan(&totalIntegral)
	if err != nil {
		log.Println("fail to select integral")
	}
	return totalIntegral
}

func GetUserInfo(redId string) (userInfo User) {
	err := DB.Raw("select * from users where redid = ?", redId).Scan(&userInfo).Error
	if err != nil {
		log.Println("fail to get userinfo",err)
	}
	return
}
