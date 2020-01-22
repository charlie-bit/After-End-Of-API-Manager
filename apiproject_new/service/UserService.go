package service

import (
	"apiproject_new/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)
/**
	Author:charlie
	Description:新增用户
	Time:2019-1-10
*/
func CreateNewUserService(user models.User) (int , map[string]interface{}) {
	godotenv.Load("../.env")
	userTest := models.User{}
	userTest = user.SelectUserByName(user.UserName)
	if userTest.UserName!= "" {
		return 500,map[string]interface{}{"message":"该用户名已注册"}
	}
	user.Ip = models.ClientIP()
	//将密码加密
	pwd,_ := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	user.Password = string(pwd)
	//jwt认证 -- 创建一个token
	var tk models.Token
	tk.UserId = user.Id
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"),tk)
	tokenString,_ := token.SignedString([]byte(os.Getenv("Token_password")))
	user.Token = tokenString
	user.EffectiveTime = time.Now().Format("2006-01-02 15:04:05")
	//失效时间为默认时间
	addTime,_ := time.ParseDuration(os.Getenv("ExpirationTime_interval"))
	user.ExpirationTime = time.Now().Add(addTime).Format("2006-01-02 15:04:05")
	//将新建的用户存储到数据库
	user.UseState = 0
	user.Role = os.Getenv("Default_Role")
	user.CreatUser()
	CreateNewUserApiRecordService(user,"Register User")
	user.Password = "" //将密码清空
	return 200,map[string]interface{}{"message":"用户创建成功","user":user}
}

/**
	Author:charlie
	Description:用户登录
	Time:2019-1-10
*/
func UserLogin(user models.User) (int,map[string]interface{}) {
	//将密码解密
	userTest := models.User{}
	userTest = user.SelectUserByName(user.UserName)
	if userTest.UserName == "" {
		return 500, map[string]interface{}{"message":"该用户名不存在"}
	}
	err := bcrypt.CompareHashAndPassword([]byte(userTest.Password),[]byte(user.Password))
	if err != nil &&  err == bcrypt.ErrMismatchedHashAndPassword{
		return 500, map[string]interface{}{"message":"密码错误"}
	}
	//判断token是否失效
	exp := userTest.ExpirationTime
	today,_ := time.ParseInLocation("2006-01-02 15:04:05",exp,time.Local)
	if today.Unix()<time.Now().Unix(){
		return 500, map[string]interface{}{"message":"token值已失效，请获得最新的token值."}
	}
	userTest.UseState = 1
	user.UpdateUser(userTest)
	CreateNewUserApiRecordService(userTest,"Login User")
	userTest.Password = ""
	return 200,map[string]interface{}{"message":"用户创建成功","user":userTest}
}
/**
	Author:charlie
	Description:获取最新的token值
	Time:2019-1-10
*/
func GetNewToken(user models.User) (int,map[string]interface{}) {
	godotenv.Load("../.env")
	userTest := models.User{}
	userTest = user.SelectUserByName(user.UserName)
	if userTest.UserName == "" {
		return 500, map[string]interface{}{"message":"该用户名不存在"}
	}
	err := bcrypt.CompareHashAndPassword([]byte(userTest.Password),[]byte(user.Password))
	if err != nil &&  err == bcrypt.ErrMismatchedHashAndPassword{
		return 500, map[string]interface{}{"message":"密码错误"}
	}
	var tk models.Token
	tk.UserId = userTest.Id
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"),tk)
	tokenString,_ := token.SignedString([]byte(os.Getenv("Token_password")))
	userTest.Token = tokenString
	userTest.EffectiveTime = time.Now().Format("2006-01-02 15:04:05")
	//失效时间为一天
	addTime,_ := time.ParseDuration(os.Getenv("ExpirationTime_interval"))
	userTest.ExpirationTime = time.Now().Add(addTime).Format("2006-01-02 15:04:05")
	user.UpdateUser(userTest)
	CreateNewUserApiRecordService(userTest,"Get New Token")
	userTest.Password = ""
	return 200,map[string]interface{}{"message":"操作成功","user":userTest}
}
