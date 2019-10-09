package model

import (
	"time"
)

type Record struct {
	ID          int `gorm:"primary_key"`
	ContentType int `gorm:"column:content_type"`
	Number      int `gorm:"column:number"`
	CreatedAt   time.Time
	
}
