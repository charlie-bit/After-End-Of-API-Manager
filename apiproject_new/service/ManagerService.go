package service

import (
	"apiproject_new/models"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"os"
)

/**
	Author:charlie
	Description:管理员登录
	Time:2019-1-14
*/
func ManagerLogin(manager models.Manager) (int,map[string]interface{}) {
	managerTest := models.Manager{}
	managerTest = manager.SelectManagerByName(manager.UserName)
	if managerTest.UserName == "" {
		return 500, map[string]interface{}{"message":"该用户名不存在"}
	}
	err := bcrypt.CompareHashAndPassword([]byte(managerTest.Password),[]byte(manager.Password))
	if err != nil &&  err == bcrypt.ErrMismatchedHashAndPassword{
		return 500, map[string]interface{}{"message":"密码错误"}
	}
	managerTest.Password = ""
	return 200,map[string]interface{}{"message":"登录成功","manager":managerTest}
}
/**
	Author:charlie
	Description:管理员注册
	Time:2019-1-14
*/
func ManagerRegister(manager models.Manager) (int,map[string]interface{}) {
	managerTest := models.Manager{}
	managerTest = manager.SelectManagerByName(manager.Password)
	if managerTest.UserName != "" {
		return 500, map[string]interface{}{"message":"该用户名已存在"}
	}
	//将密码加密
	pwd,_ := bcrypt.GenerateFromPassword([]byte(manager.Password),bcrypt.DefaultCost)
	manager.Password = string(pwd)
	manager.CreateManager()
	manager.Password = ""
	return 200,map[string]interface{}{"message":"注册成功","manager":manager}
}
/**
	Author:charlie
	Description:恢复默认密码
	Time:2019-1-14
*/
func RecoverUserPassword() (int,map[string]interface{}) {
	godotenv.Load("../.env")
	pwd,_ := bcrypt.GenerateFromPassword([]byte(os.Getenv("Default_Password")),bcrypt.DefaultCost)
	return 200,map[string]interface{}{"message":"重置成功","password":pwd}
}
