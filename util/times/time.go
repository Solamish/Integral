package times

import (
	"log"
	"mobileSign/util/https"
	"regexp"
	"strconv"
	"time"
)

var thisWeek int

func GetTimeInfo() (now, todayZero, nextZero, preZero time.Time) {
	now = time.Now().Local()
	//获得今日凌晨时间
	todayZero = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	//获得次日凌晨时间
	nextZero = todayZero.Add(time.Hour * 24)

	//获得前天凌晨时间
	preZero = todayZero.Add(-time.Hour * 24)

	return
}

func Getter() int{
	return thisWeek
}

//定时任务
func GetSchoolTime() {
	for {
		schoolTime()
		ticker := time.NewTicker(time.Hour * 24)
		<-ticker.C
	}
}

// 从教务在线获取当前教学周
func schoolTime() () {
	resp, err := https.SendGet("http://jwzx.cqu.pt/kebiao/index.php")
	if err != nil {
		log.Println("https error :", err)
	}
	html := B2S(resp)
	pattern := "第([0-9]+)周"
	match, _ := regexp.Compile(pattern)
	// 第5周
	str := match.FindAllString(html, -1)

	weekPattern := "([0-9]+)"
	weekMatch, _ := regexp.Compile(weekPattern)


	// 避免内网外入崩的时候，程序直接退出
	if len(str) <= 0 {
		log.Println("array 'str' index out of range")
		return
	}
	// 5
	week := weekMatch.FindString(str[0])

	thisWeek, _ = strconv.Atoi(week)

}


func B2S(bs []uint8) string {
	bytes := []byte{}
	for _, b := range bs {
		bytes = append(bytes, byte(b))
	}
	return string(bytes)
}

