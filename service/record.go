package service

import (
	"mobileSign/model"
)

type Record struct {
	EventType string `json:"event_type"` //type为1 表示签到;   type为2 表示提问
	Number    int    `json:"num"`
	CreatedAt string `json:"created_at"`
}

func GetRecord(redId string, page int, size int) (records []*Record) {
	var eventType string
	r := model.GetRecord(redId, page, size)

	for i, _ := range r {

		switch r[i].EventType {
		case 1:
			eventType = "签到"
		case 2:
			eventType = "提问"
		default:
			eventType = "其他"
		}

		record := Record{
			EventType: eventType,
			Number:    r[i].Number,
			CreatedAt: r[i].CreatedAt.Format("2006-01-02 15:04:05"),
		}
		records = append(records, &record)
	}
	return
}
