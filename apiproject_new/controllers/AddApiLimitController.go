package controllers

import (
	"apiproject_new/models"
	"apiproject_new/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

/**
	Author:charlie
	Description:新增一条api权限记录
	Time:2019-1-17
*/
type AddApiURLLimitStruct struct {
	ApiUrl  		string		`json:"api_url"`
	ApiDescription  string   	`json:"api_description"`
	Role 	 		string		`json:"use_role"`
}
//@Summary Add Api use rights
//@DescriptionApiManagerSystem The front end send the Request that Add Api use rights to After end
//@Tags ApiManagerSystem
//@Produce json
//@Accept json
//@Param param body controllers.AddApiURLLimitStruct true "JSON DATA FORMAT "
//@Success 200 {object} ResponseStruct
//@Failure 500 {object} ResponseStruct
//@Router /api/user/UpdateUsers [post]
func AddApiURLLimit(ctx *gin.Context)  {
	//var params UpdateUserParams
	var ps AddApiURLLimitStruct
	if err := ctx.ShouldBindJSON(&ps); err!= nil{
		utils.Respond(500,utils.Message("json数据格式不符合要求"),ctx)
		return
	}
	godotenv.Load("../.env")
	if ps.Role == "" {
		ps.Role = os.Getenv("Manager_Login_Role")
	}
	var aul models.ApiUrlLimit
	aul = aul.SelectApiLimitByApi(ps.ApiUrl)
	if aul.Id != 0 {
		utils.Respond(500,utils.Message("api路径重复"),ctx)
		return
	}
	aul.ApiUrl = ps.ApiUrl
	aul.Role = ps.Role
	aul.ApiDescription = ps.ApiDescription
	aul.AddApiLimit()
	utils.Respond(200, map[string]interface{}{"auls":aul,"message":"新增成功"},ctx)
}
/**
	Author:charlie
	Description:更新一条api权限记录
	Time:2019-1-17
*/
//@Summary Update Api use rights
//@DescriptionApiManagerSystem The front end send the Request that Update Api use rights to After end
//@Tags ApiManagerSystem
//@Produce json
//@Accept json
//@Param role query string true "角色"
//@Param description query string true "API描述"
//@Success 200 {object} ResponseStruct
//@Failure 500 {object} ResponseStruct
//@Router /api/user/UpdateApiLimit [get]
func UpdateApiLimit(ctx *gin.Context)  {
	role := ctx.Query("role")
	des := ctx.Query("description")
	id,_ := strconv.Atoi(ctx.Query("id"))
	var aul models.ApiUrlLimit
	aul.SelectApiLimit(id)
	aul.UpdateRoleLimit(role,des)
	utils.Respond(200, map[string]interface{}{"message":"更新成功"},ctx)
}
