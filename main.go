package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mobileSign/model"
	"mobileSign/route"
	"mobileSign/util/times"
)

/**
 * @Author: Solamish
 * @Description:
 * @Date: 2019/10/7 13:25
 */

func main() {
	db, err := model.InitDB()
	if err != nil {
		log.Println("err open databases", err)
		return
	}
	defer db.Close()

	go times.GetSchoolTime()
	r := gin.New()
	route.Load(r)
	r.Run(":8080")


}
