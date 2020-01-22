package models

import (
	"apiproject_new/utils"
	"regexp"
)

type Manager struct {
	Id 				int 		`gorm:"primary_key;column:manager_id";json:"manager_id"`
	UserName 		string		`gorm:"primary_key;column:manager_username";json:"username"`
	Password 		string		`gorm:"primary_key;column:manager_password";json:"password"`
}

/**
	Author:charlie
	Description:验证用户，密码
				用户名：2个英文字母以上
				密码：6位数的密码
	Time:2019-12-31
*/
func (manager *Manager)Validate() (map[string]interface{}, bool) {

	if match,_ := regexp.MatchString("^[a-zA-Z]{2,10}$",manager.UserName); !match {
		return utils.Message("用户名格式不正确"),false
	}

	if match, _ := regexp.MatchString("^\\d{6}$",manager.Password); !match {
		return utils.Message("密码格式不正确"),false
	}

	return utils.Message("该用户名，密码均可用"), true
}

/**
	Author:charlie
	Description:根据用户名查询对象
	Time:2019-1-15
*/
func (manager *Manager)SelectManagerByName(name string) Manager {
	var managerTest Manager
	GetDB().Table("managers").Where("manager_username = ?",manager.UserName).First(&managerTest)
	return managerTest
}

/**
	Author:charlie
	Description:新增对象
	Time:2019-1-10
*/
func (manager *Manager)CreateManager() {
	GetDB().Create(manager)
}


