package app

import (
	"apiproject_new/models"
	"apiproject_new/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

/**
	Author:charlie
	Description:限制有效api用户除了注册之外其他api的使用次数
	Time:2019-1-10
*/
var cache = models.GetCache()

type IpDate struct {
	Id 				int
	IpName 			string
	MaxUseNum 		int
	LastTime 		string
}

func LimitByCache(ip string) gin.HandlerFunc {
	return func(context *gin.Context) {
		//先查询一遍缓存里面有没有
		var cd IpDate
		godotenv.Load("../.env")
		//获得user上一次使用时间
		response := make(map[string]interface{})
		if cache.IsExist(ip) {
			//读取最晚的时间，可允许操作次数
			cd = cache.Get(ip).(IpDate)
			lt, _ := time.ParseInLocation("2006-01-02 15:04:05", cd.LastTime, time.Local)
			//时间间隔一分钟
			jiange, _ := time.ParseDuration("1m")
			res := time.Now().Unix() - lt.Add(jiange).Unix()
			//如果时间未超过了且请求的次数已经超过了最大限制
			//更新数据库
			if res <= 0 && cd.MaxUseNum <= 0 {
				response = utils.Message("调用太频繁  "+"截止时间:"+
					lt.Add(jiange).Format("2006-01-02 15:04:05")+"  "+
					"一分钟可以调用次数:"+os.Getenv("MaxUseNum")+"  "+
					"已调用次数:"+strconv.Itoa(cd.MaxUseNum))
				utils.Respond(429, response,context)
			}
			cd.MaxUseNum--
			cd.LastTime = time.Now().Format("2006-01-02 15:04:05")
			cache.Put(ip,cd,60*time.Second)
		}
		//将最新的时间操作api存入缓存中
		//如果时间未超过了且请求的次数已经超过了最大限制
		//更新数据库
		cd.IpName = ip
		cd.MaxUseNum,_ = strconv.Atoi(os.Getenv("MaxUseNum"))
		cd.LastTime = time.Now().Format("2006-01-02 15:04:05")
		cd.MaxUseNum--
		//将对象存入缓存
		cache.Put(ip,cd,60*time.Second)
	}
}
