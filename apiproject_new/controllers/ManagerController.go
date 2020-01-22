package controllers

import (
	"apiproject_new/models"
	"apiproject_new/service"
	"apiproject_new/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)
var store *sessions.CookieStore
var session *sessions.Session
/**
	Author:charlie
	Description:登录，注册参数
	Time:2019-1-10
*/
// JSONParams doc
type Param struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
/**
	Author:charlie
	Description:后台管理员登录
	Time:2019-1-14
*/
//@Summary ApiManagerSystem Login
//@Description Manager Login to use ApiManagerSystem
//@Tags ApiManagerSystem
//@Produce json
//@Accept json
//@Param param body controllers.JSONParam true "JSON DATA FORMAT "
//@Success 200 {object} ResponseStruct
//@Failure 500 {object} ResponseStruct
//@Router /api/manager/login [post]
func ManagerLogin(ctx *gin.Context)  {
	manager :=  models.Manager{}
	if err := ctx.ShouldBindJSON(&manager); err!= nil{
		utils.Respond(500,utils.Message("json数据格式不符合要求"),ctx)
		return
	}
	if resp, ok := manager.Validate(); !ok {
		utils.Respond(500,resp,ctx)
		return
	}
	code,resp := service.ManagerLogin(manager)
	store = sessions.NewCookieStore([]byte("secret"))
	session,_ = store.New(ctx.Request,"mySession")
	fmt.Printf("%v",session)
	utils.Respond(code,resp,ctx)
}
/**
	Author:charlie
	Description:后台管理员注册
	Time:2019-1-14
*/
//@Summary ApiManagerSystem Register
//@Description Manager Register to Login ApiManagerSystem
//@Tags ApiManagerSystem
//@Produce json
//@Accept json
//@Param param body controllers.JSONParam true "JSON DATA FORMAT "
//@Success 200 {object} ResponseStruct
//@Failure 500 {object} ResponseStruct
//@Router /api/manager/register [post]
func ManagerRegister(ctx *gin.Context)  {
	manager := models.Manager{}
	if err := ctx.ShouldBindJSON(&manager); err!= nil{
		utils.Respond(500,utils.Message("json数据格式不符合要求"),ctx)
		return
	}
	if resp, ok := manager.Validate(); !ok {
		utils.Respond(500,resp,ctx)
		return
	}
	code,resp := service.ManagerRegister(manager)
	utils.Respond(code,resp,ctx)
}

/**
	Author:charlie
	Description:查询api用户信息
	Time:2019-1-15
*/
//@Summary select the message corresponding databases
//@Description ApiManagerSystem The front end send the 	Request that select the message corresponding databases to After end
//@Tags ApiManagerSystem
//@Produce json
//@Accept json
//@Param page query string true "当前页"
//@Param table query string true "当前数据库表"
//@Success 200 {object} ResponseStruct
//@Failure 500 {object} ResponseStruct
//@Router /api/manager/selectManagerInfo [get]
func SelectApiUserInfo(ctx *gin.Context)  {
	if store == nil{
		utils.Respond(500, map[string]interface{}{"message":"请先登录"},ctx)
		return
	}
	fmt.Println(store)
	session,_ := store.Get(ctx.Request,"mySession")
	if session == nil{
		utils.Respond(500, map[string]interface{}{"message":"请先登录"},ctx)
		return
	}
	godotenv.Load("../.env")
	//第一次打开页面
	page,_ := strconv.Atoi(ctx.Query("page"))
	table := ctx.Query("table")
	if page == 0 {
		page = 1
	}
	pager := models.Pager{}
	pager.Page = page
	pager.Size,_ = strconv.Atoi(os.Getenv("pagesize"))
	var u interface{}
	user := models.User{}
	u = user.SelectUserAll(pager.Page,pager.Size,table) //查询的数据
	pager.Total = user.SelectUserTotal(pager.Size,table)
	//将角色查出来 方便后面使用
	var role models.Role
	var rs []models.Role
	rs = role.SelectRoleAll()
	utils.Respond(200, map[string]interface{}{"users":u,"ps":pager,"roles":rs},ctx)
}
