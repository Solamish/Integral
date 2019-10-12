package model

import (
	"database/sql"
	"log"
	"time"
)

type Record struct {
	ID        int    `gorm:"primary_key"`
	RedId     string `gorm:"column:redid"`
	EventType int    `gorm:"column:event_type"` // type为1 表示签到;   type为2 表示提问; default: 其他
	Number    int    `gorm:"column:number"`
	CreatedAt time.Time
}

func GetRecord(redId string, page, size int) (records []*Record) {
	var err error
	var rows *sql.Rows
	if size > 0 {
		rows, err = DB.Raw("select * from records where redid = ? order by created_at  desc limit ? offset ?", redId, size, (size-1)*page).Rows()
	} else {
		rows, err = DB.Raw("select * from records where redid = ? order by created_at  desc limit 6 offset 1", redId).Rows()
	}
	if err != nil {
		log.Println("fail to get user's record", err)
	}
	defer rows.Close()

	for rows.Next() {
		var record Record
		scanErr := DB.ScanRows(rows, &record)
		if scanErr != nil {
			log.Println("fail to scan record to struct", scanErr)
		}
		records = append(records, &record)
	}
	return
}

func (record *Record) AddRecord() {
	err := DB.Create(record).Error
	if err != nil {
		log.Println("fail to insert a record", err)
	}
}
