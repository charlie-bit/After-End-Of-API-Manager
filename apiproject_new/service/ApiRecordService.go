package service

import (
	"apiproject_new/models"
	"strconv"
	"time"
)
/**
	Author:charlie
	Description:新增api使用记录
	Time:2019-1-10
*/
func CreateNewUserApiRecordService(user models.User,string2 string)  {
	var apirecord models.ApiRecord
	startTime := time.Now()
	apirecord.Ip = user.Ip
	apirecord.UserName = user.UserName
	apirecord.ApiEndPoint = string2
	apirecord.StartTime = startTime.Format("2006-01-02 15:04:05")
	apirecord.State = 200
	apirecord.Message = string2 +" successfully"
	endTime := time.Now()
	apirecord.EndTime = endTime.Format("2006-01-02 15:04:05")
	res := endTime.Sub(startTime)
	sub := res.Milliseconds()
	apirecord.ResponseTime = strconv.FormatFloat(float64(sub),'f',-1,64)+"ms"
	apirecord.ApiEndPoint = "Register user"
	apirecord.CreateNewApiRecord()
}
