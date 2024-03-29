package service

import (
	"mobileSign/model"
	"mobileSign/util/times"
	"time"
)

var (
	contDays        = 0
	additionalScore = 0
	baseScore       = 10
	lastSignTime    time.Time
)

func Sign(stuNum string) (int) {

	now, todayZero, nextZero, preZero := times.GetTimeInfo()
	lastSignTime = model.GetLastSignTime(stuNum)

	thisWeek := times.Getter()

	if thisWeek < 0 || thisWeek > 20 {
		//fmt.Println("不在签到日期内")

		return -2
	}

	//上次签到时间在今日零点和次日零点之间
	if todayZero.Before(lastSignTime) && lastSignTime.Before(nextZero) {
		// TODO 签到后的处理
		//fmt.Println("签到过了")

		return -1
	}

	// 如果是周一，签到天数重置
	if now.Weekday() == time.Monday {
		model.ResetContDays(stuNum, now)

		//fmt.Println("周一签到")
	} else {
		//如果不是周一，且为连续签到
		if lastSignTime.Before(todayZero) && lastSignTime.After(preZero) {
			//连续签到天数+1
			//fmt.Println("连续签到")
			model.UpdateContDays(stuNum, now)

			contDays = model.GetContDays(stuNum)

			switch contDays {
			case 2:
				additionalScore = 5
			case 3:
				additionalScore = 10
			case 4:
				additionalScore = 15
			case 5:
				additionalScore = 20
			case 6:
				additionalScore = 20
			case 7:
				additionalScore = 20
			}

		} else {
			//fmt.Println("第一次签到")
			model.ResetContDays(stuNum, now)

		}
	}

	score := baseScore + additionalScore
	record := model.Record{
		StuNum:    stuNum,
		EventType: 1,
		Number:    score,
	}

	// 增加积分
	model.UpdateIntegral(stuNum, score)

	// 添加一条积分纪录
	record.AddRecord()
	return score
}
