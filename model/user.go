package model

import (
	"log"
	"time"
)

type User struct {
	ID                int       `gorm:"primary_key"`
	RedId             string    `gorm:"column:redid"`
	Stunum            string    `gorm:"column:stu_num"`
	Introduction      string    `gorm:"column:introduction"`
	UserName          string    `gorm:"column:username"`
	NickName          string    `gorm:"column:nickname"`
	Gendor            string    `gorm:"column:gendor"`
	HeadImgUrl        string    `gorm:"column:head_img_url"`
	LastSignTime      time.Time `gorm:"column:last_sign_time"`
	PhotoSrc 		  string  	`gorm:"column:photo_src"`
	PhotoThumbnailSrc string    `gorm:"column:photo_thumbnail_src"`   // 缩略图
	Phone             string    `gorm:"column:phone"`
	QQ                string    `gorm:"column:qq"`
	Integral          int       `gorm:"column:integral"`
	CheckInDays       int       `gorm:"column:check_in_days"`
}

func (user *User) Save() {
	tmp := User{}
	DB.Where("stu_num = ?", user.Stunum).Find(&tmp)
	if tmp.Stunum == "" {
		DB.Create(user)
	}
}

func GetLastSignTime(stuNum string) (lastSignTime time.Time) {
	err := DB.Raw("select last_sign_time from users where stu_num = ?", stuNum).Row().Scan(&lastSignTime)
	if err != nil {
		log.Println("fail to get last sign time", err)
	}
	return
}

func GetContDays(stuNum string) (contDays int) {
	err := DB.Raw("select check_in_days from users where stu_num = ?", stuNum).Row().Scan(&contDays)
	if err != nil {
		log.Println("fail to get check_in_days", err)
	}
	return
}

func ResetContDays(stuNum string, now time.Time) {
	//err := DB.Table("users").Where("redid = ?", redId).Updates(map[string]interface{}{"cont_days": 1, "last_sign_time": now,}).Error
	err := DB.Exec("UPDATE users SET check_in_days = 1, last_sign_time = ?  WHERE stu_num = ?", now, stuNum).Error
	if err != nil {
		log.Println("fail to reset check_in_days days", err)
	}

}

func UpdateContDays(stuNum string, now time.Time) {
	err := DB.Exec("UPDATE users SET check_in_days = check_in_days + 1, last_sign_time = ? WHERE stu_num = ?", now, stuNum).Error
	if err != nil {
		log.Println("fail to update check_in_days days", err)
	}
}

func UpdateIntegral(stuNum string, integral int) {
	err := DB.Exec("UPDATE users SET integral = integral + ? WHERE stu_num = ?", integral, stuNum).Error
	if err != nil {
		log.Println("fail to update integral", err)
	}
}

func GetIntegral(stuNum string) (totalIntegral int) {
	err := DB.Raw("select integral from users where stu_num = ?", stuNum).Row().Scan(&totalIntegral)
	if err != nil {
		log.Println("fail to select integral")
	}
	return totalIntegral
}

func GetUserInfo(stuNum string) (userInfo User) {
	err := DB.Raw("select * from users where stu_num = ?", stuNum).Scan(&userInfo).Error
	if err != nil {
		log.Println("fail to get userinfo",err)
	}
	return
}
