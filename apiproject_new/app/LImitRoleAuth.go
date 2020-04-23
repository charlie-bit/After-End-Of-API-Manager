package app

import (
	"After-End-Of-API-Manager/apiproject_new/models"
	"After-End-Of-API-Manager/apiproject_new/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

/**
Author:charlie
Description:角色验证
Time:2019-12-31
*/
var LimitRoleAuthentication = func() gin.HandlerFunc {

	return gin.HandlerFunc(func(ctx *gin.Context) {
		//当前路径
		requestPath := ctx.Request.URL.Path
		response := make(map[string]interface{})
		//可以跳过验证权限的apiurl
		notAuth := []string{"/api/user/new", "/api/user/getToken", "/api/user/login"}
		for _, val := range notAuth {
			if val == requestPath {
				ctx.Next()
				return
			}
		}
		/**
		验证身份
		首先：验证身份是否存在
		其次：判断是否有权限
		*/
		name := ctx.Request.URL.Query().Get("username")
		//用名字查是否登录过
		var user models.User
		err := models.GetDB().Where("user_name=?", name).First(&user).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			response = utils.Message("数据库连接错误")
			utils.Respond(500, response, ctx)
			return
		}
		if err == gorm.ErrRecordNotFound {
			response = utils.Message("该账户不存在")
			utils.Respond(500, response, ctx)
			return
		}
		if user.UseState != 1 {
			response = utils.Message("该账户还未登录！！")
			utils.Respond(500, response, ctx)
			return
		}
		var aul models.ApiUrlLimit
		err = models.GetDB().Where("use_role=? and api_url=?", user.Role, requestPath).First(&aul).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			response = utils.Message("数据库连接错误")
			utils.Respond(500, response, ctx)
			return
		}
		if err == gorm.ErrRecordNotFound {
			response = utils.Message("还用户无权限")
			utils.Respond(500, response, ctx)
			return
		}
		ctx.Next()
	})
}
