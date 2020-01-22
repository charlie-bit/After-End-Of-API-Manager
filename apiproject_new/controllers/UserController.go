package controllers

import (
	"apiproject_new/app"
	"apiproject_new/models"
	"apiproject_new/service"
	"apiproject_new/utils"
	"github.com/gin-gonic/gin"
	"regexp"
	"strconv"
)

/**
	Author:charlie
	Description:新增用户
	Time:2019-1-10
*/
// JSONParams doc
type JSONParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
// JSONParams doc
//@Summary NewUser
//@Description New user to Login
//@Tags User
//@Produce json
//@Accept json
//@Param param body controllers.JSONParam true "JSON DATA FORMAT "
//@Success 200 {object} ResponseStruct
//@Failure 500 {object} ResponseStruct
//@Router /api/user/new [post]
func NewUser(ctx *gin.Context)  {
	params := &JSONParam{}
	if err := ctx.BindJSON(params); err!= nil{
		utils.Respond(500,utils.Message("json数据格式不符合要求"),ctx)
		return
	}
	user :=  models.User{}
	user.UserName = params.Username
	user.Password = params.Password
	if resp, ok := user.Validate(); !ok {
		utils.Respond(500,resp,ctx)
		return
	}
	code,resp := service.CreateNewUserService(user)
	var respstruct ResponseStruct
	respstruct.Code = code
	respstruct.Message = resp
	utils.Respond(code,resp,ctx)
}

/**
	Author:charlie
	Description:用户登录
	Time:2019-1-10
*/
//@Summary LoginUser
//@Description Login user to user api
//@Tags User
//@Produce json
//@Accept json
//@Param param body controllers.JSONParam true "JSON DATA FORMAT "
//@Success 200 {object} ResponseStruct
//@Failure 500 {object} ResponseStruct
//@Router /api/user/login [post]
//@Security BearerTokenAuth
func LoginUser(ctx *gin.Context)  {
	if ok :=app.JwtAuthentication(ctx);!ok{
		return
	}
	user :=  models.User{}
	if err := ctx.ShouldBindJSON(&user); err!= nil{
		utils.Respond(500,utils.Message("json数据格式不符合要求"),ctx)
		return
	}
	if resp, ok := user.Validate(); !ok {
		utils.Respond(500,resp,ctx)
	}
	code,resp := service.UserLogin(user)
	utils.Respond(code,resp,ctx)
}

/**
	Author:charlies
	Description:获取最新的token值
	Time:2019-1-10
*/
//@Summary Get New Token
//@Description If the EffectTime out the ExpirationTime, Please Get the new Token
//@Tags User
//@Produce json
//@Accept json
//@Param param body controllers.JSONParam true "JSON DATA FORMAT "
//@Success 200 {object} ResponseStruct
//@Failure 500 {object} ResponseStruct
//@Router /api/user/getToken [post]
func GetNewToken(ctx *gin.Context)  {
	user :=  models.User{}
	if err := ctx.ShouldBindJSON(&user); err!= nil{
		utils.Respond(500,utils.Message("json数据格式不符合要求"),ctx)
		return
	}
	if resp, ok := user.Validate(); !ok {
		utils.Respond(500,resp,ctx)
	}
	code,resp := service.GetNewToken(user)
	utils.Respond(code,resp,ctx)
}

/**
	Author:charlies
	Description:根据前台提交过来的数据进行数据库更新
	Time:2019-1-17
*/
type UpdateUserParams struct {
	Id 				int 	`json:"Id"`
	Name 			string 	`json:"username"`
	Password 		string 	`json:"password"`
	Expiration_time string 	`json:"expiration_time"`
	Use_role        string 	`json:"use_role"`
}
//@Summary Update the message of the User using Api
//@DescriptionApiManagerSystem The front end send the Request that Update the message of the User using Api to After end
//@Tags ApiManagerSystem
//@Produce json
//@Accept json
//@Param param body controllers.UpdateUserParams true "JSON DATA FORMAT "
//@Success 200 {object} ResponseStruct
//@Failure 500 {object} ResponseStruct
//@Router /api/user/UpdateUsers [post]
func UpdateUsers(ctx *gin.Context)  {
	//var params UpdateUserParams
	var params UpdateUserParams
	if err := ctx.ShouldBindJSON(&params); err!= nil{
		utils.Respond(500,utils.Message("json数据格式不符合要求"),ctx)
		return
	}
	var user models.User
	if match,_ := regexp.MatchString("^[a-zA-Z]{2,}$",params.Name); !match {
		utils.Respond(500, map[string]interface{}{"messages":"用户名格式不正确"},ctx)
		return
	}
	user = user.SelectUserByName(params.Name)
	if user.UserName!= "" {
		utils.Respond(500, map[string]interface{}{"messages":"用户名已存在"},ctx)
		return
	}
	user = user.SelectUserById(params.Id)
	user.UpdateUserByMap(map[string]interface{}{"user_name":params.Name,"user_password":params.Password,
		"expiration_time":params.Expiration_time,"use_role":params.Use_role})
	utils.Respond(200, map[string]interface{}{"users":user,"message":"更新成功"},ctx)
}

/**
	Author:charlie
	Description:恢复默认密码
	Time:2019-1-17
*/
//@Summary Recover default password corresponding button
//@DescriptionApiManagerSystem The front end send the Request that recover the users' default password to After end
//@Tags ApiManagerSystem
//@Produce json
//@Accept json
//@Success 200 {object} ResponseStruct
//@Failure 500 {object} ResponseStruct
//@Router /api/user/RecorverPassword [post]
func RecoverPassword(ctx *gin.Context)  {
	code,resp := service.RecoverUserPassword()
	utils.Respond(code,resp,ctx)
}

/**
	Author:charlie
	Description:删除指定记录
	Time:2019-1-17
*/
//@Summary select the message corresponding databases
//@Description ApiManagerSystem The front end send the 	Request that select the message corresponding databases to After end
//@Tags ApiManagerSystem
//@Produce json
//@Accept json
//@Param id query integer true "指定id"
//@Param table query string true "当前数据库表"
//@Success 200 {object} ResponseStruct
//@Failure 500 {object} ResponseStruct
//@Router /api/manager/Delete [get]
func DeleteUser(ctx *gin.Context)  {
	id,_ := strconv.Atoi(ctx.Query("id"))
	table := ctx.Query("table")
	if table == "users" {
		user := models.User{}
		user.DeleteUserById(id)
	}
	if table == "api_url_limits" {
		aul := models.ApiUrlLimit{}
		aul.DeleteApiUrlLimitsById(id)
	}
	utils.Respond(200, map[string]interface{}{"message":"删除成功"},ctx)
}
